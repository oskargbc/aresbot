package shops

import (
	models2 "aresbot/internal/aio/models"
	"aresbot/internal/aio/shared"
	l "aresbot/pkg/logger"
	"fmt"
	"sync"
	"time"
)

type BotRunner struct {
	Tasks          []models2.Task
	Proxys         []models2.Proxy
	Settings       models2.Settings
	WebhookHandler shared.WebhookHandler
}

func (b *BotRunner) Run() {
	var wg sync.WaitGroup
	rotator := shared.ProxyRotator{Proxys: b.Proxys}

	for id, task := range b.Tasks {
		shop := getShopByTask(task)
		id++

		//fmt.Println(&wg, shop, task, b.Settings, id)
		go ShopRunner(&wg, shop, task, b.Settings, id, rotator, b.WebhookHandler)

		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("")
}

func ShopRunner(wg *sync.WaitGroup, bot ShopBot, task models2.Task, settings models2.Settings, id int, r shared.ProxyRotator, sender shared.WebhookHandler) {
	var foundedProduct models2.Product

	if !bot.IsActive() {
		shared.OError(id, "Module locked.")
		wg.Done()
		return
	}

	shared.OInfo(id, "Starting... ")
	l.InfoLogger.Println("Starting ", task, "with bot ", &bot)

	task.Id = id

	err := bot.GetSettings(settings)
	if err != nil {
		l.ErrorLogger.Println(err)
		shared.OError(id, err.Error())
		wg.Done()
		return
	}

	var proxy *models2.Proxy

	if task.UseProxy {
		proxy = r.RotateProxy()

		if proxy == nil {
			l.ErrorLogger.Println("Couldn't find proxy ", task)
			shared.OError(id, "Couldn't find Proxy")

			wg.Done()
			return
		}
	}
	// Waiting for product
	shared.OInfo(id, "Waiting for Product-Drop..")

	retry := settings.RetryCount
	if retry == 0 {
		retry = 1
	}

	var product *models2.Product
	success := false

	// Monitoring Products
	for retry > 0 {
		resp := bot.CheckProduct(task, proxy)

		if resp.rotateProxy {
			shared.OInfo(id, "Rotating Proxy..")
			proxy = r.RotateProxy()

			if proxy == nil {
				l.ErrorLogger.Println("Couldn't find proxy ", task)
				shared.OError(id, "Couldn't rotate Proxy, stopping task")

				wg.Done()
				return
			}
		}

		if resp.e != nil {
			shared.OWarning(id, "Error while searching Product, trying again")
			l.ErrorLogger.Println(resp.e)
			retry--
			continue
		}

		if resp.success && resp.product != nil {
			foundedProduct = *resp.product
			l.InfoLogger.Println(foundedProduct)

			success = true
			retry = -1
		} else {
			mdelay := settings.MonitorDelay
			time.Sleep(time.Duration(mdelay) * time.Millisecond)
		}
	}

	if &product == nil || !success {
		shared.OError(id, "Failed to find product.. stopping task")

		wg.Done()
		return
	}

	shared.OInfo(id, "Found Product..")

	if bot.NeedCookies(task) {
		// First Step, generate Cookies
		shared.OInfo(id, "Generating Cookies...")
		l.InfoLogger.Println(task, "Generating Cookies")
		retryCount := settings.RetryCount
		gotCookies := false

		for retryCount > 0 {
			resp := bot.GenerateCookies(task, proxy)

			// Error handling
			// if there is an error
			if resp.e != nil || !resp.success {
				l.ErrorLogger.Println(resp.e)

				if resp.rotateProxy {
					shared.OInfo(id, "Rotating Proxy..")
					proxy = r.RotateProxy()

					if proxy == nil {
						l.ErrorLogger.Println("Couldn't find proxy ", task)
						shared.OError(id, "Couldn't rotate Proxy, stopping task")

						wg.Done()
						return
					}
				}
				time.Sleep(time.Second * 2)
				retryCount--
				continue
			}

			// rotate Proxy
			if resp.rotateProxy {
				shared.OInfo(id, "Rotating Proxy..")
				proxy = r.RotateProxy()

				if proxy == nil {
					l.ErrorLogger.Println("Couldn't find proxy ", task)
					shared.OError(id, "Couldn't rotate Proxy, stopping task")

					wg.Done()
					return
				}
			}

			// if cookies are need and got generated
			if resp.success {
				retryCount = -1
				gotCookies = true
				shared.OStep(id, "Generated Cookies")
			}
		}
		if !gotCookies {
			shared.OError(id, "Couldn't generate Cookies")

			wg.Done()
			return
		}
	}

	if bot.NeedLogin(task) {
		shared.OInfo(id, "Logging in..")

		sLogin := false

		retryLogin := settings.RetryCount

		if retryLogin == 0 {
			retryLogin = 1
		}

		for retryLogin > 0 {
			// Trying Login
			resp := bot.Login(task, proxy)

			if resp.e != nil {
				l.ErrorLogger.Println(resp.e, task)
				shared.OError(id, "Logging in failed, trying again")

				retryLogin--
			}

			if resp.rotateProxy {
				shared.OInfo(id, "Rotating Proxy..")
				proxy = r.RotateProxy()

				if proxy == nil {
					l.ErrorLogger.Println("Couldn't find proxy ", task)
					shared.OError(id, "Couldn't rotate Proxy, stopping task")

					wg.Done()
					return
				}
			}

			if resp.success {
				sLogin = true
				retryLogin = -1
			} else {
				l.ErrorLogger.Println("Failed login", task)
				shared.OWarning(id, "Failed Login, trying again..")
				retryLogin--
			}
		}

		if !sLogin {
			shared.OError(id, "Login failed, stopping task")

			wg.Done()
			return
		} else {
			shared.OStep(id, "Login succeed")
		}
	}

	// Starting Atc
	hasAtc := false

	retryAtc := settings.RetryCount
	if retryAtc == 0 {
		retryAtc = 1
	}
	shared.OInfo(id, "Trying add to cart..")

	for retryAtc > 0 {
		resp := bot.AddToCart(task, proxy)

		if resp.rotateProxy {
			shared.OInfo(id, "Rotating Proxy..")
			proxy = r.RotateProxy()

			if proxy == nil {
				l.ErrorLogger.Println("Couldn't find proxy ", task)
				shared.OError(id, "Couldn't rotate Proxy, stopping task")

				wg.Done()
				return
			}
		}

		if resp.e != nil {
			shared.OWarning(id, "Failed Adding to Cart, trying again...")
			l.ErrorLogger.Println(resp.e, task)
			retryAtc--
			continue
		}

		if resp.success {
			l.InfoLogger.Println(task, product, "Added to cart")

			hasAtc = true
			retryAtc = -1
		}
	}

	if !hasAtc {
		shared.OError(id, "Add To Cart failed.. stopping task")

		wg.Done()
		return
	} else {
		shared.OStep(id, "Product added to cart")
	}

	// Submit Shipping
	hasShipping := false

	retryShipping := settings.RetryCount
	if retryShipping == 0 {
		retryShipping = 1
	}
	shared.OInfo(id, "Submit Shipping..")

	for retryShipping > 0 {
		resp := bot.SubmitShipping(task, proxy)

		if resp.rotateProxy {
			shared.OInfo(id, "Rotating Proxy..")
			proxy = r.RotateProxy()

			if proxy == nil {
				l.ErrorLogger.Println("Couldn't find proxy ", task)
				shared.OError(id, "Couldn't rotate Proxy, stopping task")

				wg.Done()
				return
			}
		}

		if resp.e != nil {
			shared.OWarning(id, "Failed Submit Shipping, trying again...")
			l.ErrorLogger.Println(resp.e, task)
			retryShipping--
			continue
		}

		if resp.success {
			l.InfoLogger.Println(task, product, "Submitted Shipping")

			hasShipping = true
			retryShipping = -1
		}
	}

	if !hasShipping {
		shared.OError(id, "Submitting Shipping failed.. stopping task")

		wg.Done()
		return
	} else {
		shared.OStep(id, "Submitted Shipping")
	}

	// Submit Billing
	hasBilling := false

	retryBilling := settings.RetryCount
	if retryBilling == 0 {
		retryBilling = 1
	}
	shared.OInfo(id, "Submit Billing..")

	for retryBilling > 0 {
		resp := bot.SubmitBilling(task, proxy)

		if resp.rotateProxy {
			shared.OInfo(id, "Rotating Proxy..")
			proxy = r.RotateProxy()

			if proxy == nil {
				l.ErrorLogger.Println("Couldn't find proxy ", task)
				shared.OError(id, "Couldn't rotate Proxy, stopping task")

				wg.Done()
				return
			}
		}

		if resp.e != nil {
			shared.OWarning(id, "Failed Submit Billing, trying again...")
			l.ErrorLogger.Println(resp.e, task)
			retryBilling--
			continue
		}

		if resp.success {
			l.InfoLogger.Println(task, product, "Submitted Billing")

			hasBilling = true
			retryBilling = -1
		}
	}

	if !hasBilling {
		shared.OError(id, "Billing failed.. stopping task")

		wg.Done()
		return
	} else {
		shared.OStep(id, "Submitted Billing")
	}

	// Starting Checkout
	shared.OInfo(id, "Trying Checkout..")

	resp := bot.Checkout(task, proxy, foundedProduct)

	if resp.e != nil {
		shared.OError(id, "Failed Checkout with Error")
		l.ErrorLogger.Println(resp.e, task, "Checkout Failed with Error")

		wg.Done()
		return
	}

	if resp.success {
		l.InfoLogger.Println(task, product, "Checkout Success")
		shared.OSuccess(id, "Successful Checkout")

		go sender.SendWebhook(shared.WebhookData{
			Product: *resp.product,
			Profile: task.Profile,
		})

		wg.Done()
		return
	} else {
		l.ErrorLogger.Println(task, "Checkout Failed, OOS")
		shared.OError(id, "Checkout Failed, OOS")

		wg.Done()
		return
	}
}

func getShopByTask(task models2.Task) ShopBot {

	return ShopBots[task.Store]
}

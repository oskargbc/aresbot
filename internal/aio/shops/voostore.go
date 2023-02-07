package shops

import (
	"aresbot/internal/aio/constants"
	"aresbot/internal/aio/models"
	"aresbot/internal/aio/shared"
	"aresbot/internal/aio/shops/einhalb"
	"aresbot/internal/aio/shops/voostore"
	"aresbot/pkg/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

func NewVoostoreBot() ShopBot {
	return &Voostore{}
}

const (
	AddToCartUrl    = "https://www.vooberlin.com/checkout/ajaxAddArticleCart"
	CheckoutUrl     = "https://www.vooberlin.com/checkout/ajaxAddArticleCart"
	CreatePaypalUrl = "https://www.vooberlin.com/widgets/PaypalUnifiedExpressCheckout/createPayment"
	PaypalTokenUrl  = "https://www.paypal.com/smart/api/payment/%s/ectoken"
	GetSizeBaseUrl  = "%s?group[1]=%s&template=ajax"
)

type Voostore struct {
	ShopBot

	Client    *http.Client
	SizeAndId map[string]string
	PaymentId string
}

func (s *Voostore) CheckProduct(task models.Task, proxy *models.Proxy) Response {
	if proxy != nil {
		s.Client.Transport = &http.Transport{Proxy: http.ProxyURL(&proxy.Url)}
	}

	sizeAndId := map[string]string{}

	resp, err := s.Client.Get(task.ProductURL)
	if err != nil {
		return Response{e: err}
	}

	if resp.StatusCode == 200 {
		s.Client.Jar.SetCookies(resp.Request.URL, resp.Cookies())

		bod, _ := ioutil.ReadAll(resp.Body)

		doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(bod))

		doc.Find("body > div.page-wrap > section > div > div.content--wrapper > div > div.product--detail-upper.block-group > div.product--buybox.block > div > div.product--configurator > form > div > select > option").Each(func(i int, selection *goquery.Selection) {
			size := strings.TrimSpace(selection.Text())
			id, ok := selection.Attr("value")
			if ok {
				sizeAndId[size] = id
			}
		})

		img := ""
		mg, ok := doc.Find("body > div.page-wrap > section > div > div.content--wrapper > div > div.product--detail-upper.block-group > div.product--image-container.image-slider > div.image-slider--container > div > div:nth-child(1) > span > span > picture > img").Attr("srcset")
		if ok {
			img = mg
		}
		name := doc.Find("body > div.page-wrap > section > div > div.content--wrapper > div > div.product--detail-upper.block-group > header > div.product--info > h1 > span").Text()
		if len(task.Size) > 6 {
			task.Sku = task.Size
		}
		if len(task.Sku) > 5 {
			return Response{success: true, product: &models.Product{
				Name:     name,
				Size:     task.Sku,
				Url:      task.ProductURL,
				ImageUrl: img,
				Payment:  task.Payment,
				Store:    task.Store,
			}}
		}
		doc.Find("body > div.page-wrap > section > div > div.content--wrapper > div > div.product--detail-upper.block-group > div.product--buybox.block > div > div.product--configurator > form > div > select > option").Each(func(i int, selection *goquery.Selection) {
			size := strings.TrimSpace(selection.Text())
			id, ok := selection.Attr("value")
			if ok {
				sizeAndId[size] = id
			}
		})
		for k, _ := range sizeAndId {
			if k == task.Size {
				return Response{success: true, product: &models.Product{
					Name:     name,
					Size:     task.Size,
					Url:      task.ProductURL,
					ImageUrl: img,
					Payment:  task.Payment,
					Store:    task.Store,
				}}
			}
		}
		l := 0
		for _, _ = range sizeAndId {
			l++
		}
		i := rand.Intn(l)
		j := 0
		if task.Size != "RANDOM" {
			shared.OInfo(task.Id, "size not in stock, choosing random")
		}

		for k, _ := range sizeAndId {
			if j == i {
				task.Size = k
				return Response{success: true, product: &models.Product{
					Name:     name,
					Size:     task.Size,
					Url:      task.ProductURL,
					ImageUrl: img,
					Payment:  task.Payment,
					Store:    task.Store,
				}}
			}
			j++
		}

		s.SizeAndId = sizeAndId
		return Response{}
	} else if resp.StatusCode == 403 {
		logger.WarningLogger.Println(resp.StatusCode)

		return Response{success: false, rotateProxy: true}
	} else if resp.StatusCode == 429 {
		logger.WarningLogger.Println(resp.StatusCode)

		return Response{success: false, rotateProxy: true}
	} else {
		logger.ErrorLogger.Println(resp.StatusCode)

		return Response{e: constants.ErrUnexpectedStatusCode, success: false}
	}
}
func (s *Voostore) GenerateCookies(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}
func (s *Voostore) Login(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}
func (s *Voostore) AddToCart(task models.Task, proxy *models.Proxy) Response {
	shared.OInfo(task.Id, "Waiting for product to be cartable")
	for {
		if proxy != nil {
			s.Client.Transport = &http.Transport{Proxy: http.ProxyURL(&proxy.Url)}
		}

		if len(task.Sku) < 5 {
			task.Sku = s.GetSku(task.ProductURL, s.SizeAndId[task.Size])
		}

		if task.Sku == "" {
			return Response{e: constants.NewGeneralError("sku not found")}
		}

		p := "sActionIdentifier=&sAddAccessories=&sAdd=" + task.Sku + "&sQuantity=1"
		payload := strings.NewReader(p)

		req, _ := http.NewRequest(http.MethodPost, AddToCartUrl, payload)

		for k, v := range voostore.VooStoreAddToCartHeaders {
			req.Header.Add(k, v)
		}
		req.Header.Add("referer", task.ProductURL)

		resp, err := s.Client.Do(req)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return Response{
				e:       err,
				success: false,
			}
		}
		if resp.StatusCode != 200 {
			logger.WarningLogger.Println(resp.StatusCode)
			return Response{e: constants.ErrUnexpectedStatusCode}
		}

		if strings.Contains(string(body), "The product was successfully added to your shopping cart") {
			s.Client.Jar.SetCookies(resp.Request.URL, resp.Cookies())

			return Response{success: true}
		}
		if strings.Contains(string(body), "This product cannot be purchased at the moment!") {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		logger.ErrorLogger.Println(resp.StatusCode)
		return Response{e: constants.ErrAtcStepFailed}
	}
}

func (s *Voostore) SubmitBilling(task models.Task, proxy *models.Proxy) Response {
	return Response{success: true}
}

func (s *Voostore) SubmitShipping(task models.Task, proxy *models.Proxy) Response {
	if proxy != nil {
		s.Client.Transport = &http.Transport{Proxy: http.ProxyURL(&proxy.Url)}
	}

	payload := strings.NewReader(`useInContext=true`)

	req, _ := http.NewRequest(http.MethodPost, CreatePaypalUrl, payload)

	for k, v := range voostore.VooStoreCreatePaypalHeaders {
		req.Header.Add(k, v)
	}
	req.Header.Add("referer", task.ProductURL)

	resp, err := s.Client.Do(req)
	if err != nil {
		return Response{e: err}
	}

	logger.InfoLogger.Println(resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{e: err}
	}
	var paypalResp voostore.CreatePaypalResponse
	err = json.Unmarshal(body, &paypalResp)
	if err != nil {
		return Response{e: err}
	}

	s.PaymentId = paypalResp.PaymentId
	s.Client.Jar.SetCookies(resp.Request.URL, resp.Cookies())
	return Response{success: true}
}

func (s *Voostore) Checkout(task models.Task, proxy *models.Proxy, product models.Product) Response {
	if proxy != nil {
		s.Client.Transport = &http.Transport{Proxy: http.ProxyURL(&proxy.Url)}
	}

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf(PaypalTokenUrl, s.PaymentId), strings.NewReader(`{"meta":{}}`))

	for k, v := range einhalb.PaypalCheckoutHeaders {
		req.Header.Add(k, v)
	}

	resp, err := s.Client.Do(req)

	if err != nil {
		return Response{success: false, e: err}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorLogger.Println(string(body))
		return Response{e: err}
	}

	var tokenResp einhalb.PaypalTokenResponse

	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return Response{e: err}
	}

	if tokenResp.Ack != "success" {
		logger.ErrorLogger.Println("checkout token creation error")
		return Response{success: false}
	}

	checkoutUrl := "https://www.paypal.com/checkoutnow?version=4&token=" + tokenResp.Data.Token

	product.CheckoutUrl = checkoutUrl

	return Response{success: true, product: &product}
}
func (s *Voostore) NeedLogin(task models.Task) bool {
	return false
}
func (s *Voostore) NeedCookies(task models.Task) bool {
	return false
}
func (s *Voostore) GetSettings(settings models.Settings) error {
	j, _ := cookiejar.New(nil)
	s.Client = &http.Client{Jar: j}

	return nil
}
func (s *Voostore) IsActive() bool {
	return true
}

func (s *Voostore) GetSku(productUrl, id string) string {
	uri := fmt.Sprintf(GetSizeBaseUrl, productUrl, id)

	resp, err := s.Client.Get(uri)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return ""
	}

	if resp.StatusCode != 200 {
		logger.WarningLogger.Println(resp.StatusCode)
		return ""
	}

	bod, _ := ioutil.ReadAll(resp.Body)
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(bod))
	sku := doc.Find("body > div.page-wrap > section > div > div.content--wrapper > div > div > div.product--buybox.block > ul > div > span").Text()

	return sku
}

package shared

import (
	"aresbot/internal/aio/models"
	l "aresbot/pkg/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

type WebhookHandler struct {
	Version   string
	PublicUrl string
	PrivatUrl string
	HyperData HyperResponse
}

func (w *WebhookHandler) SendWebhook(d WebhookData) {
	c := http.Client{}
	t := time.Now()

	user := "not found"
	if w.HyperData.User != nil {
		user = w.HyperData.User.(map[string]interface{})["username"].(string)
	}

	plan := w.HyperData.Plan.Name

	privat := Webhook{
		Content:   "",
		Username:  "AresBot",
		AvatarURL: nil,
		Tts:       false,
		Embeds: []Embeds{
			Embeds{
				Title:       "Finish Checkout  :fire:",
				Description: "Successful Checkout for " + IfNilEmpty(&d.Product.Name),
				URL:         IfNilPlaceholder(&d.Product.CheckoutUrl),
				Timestamp:   &t,
				Color:       850000,
				Footer: Footer{
					Text:    "AresBot " + IfNilEmpty(&w.Version),
					IconURL: nil,
				},
				Image: Image{},
				Thumbnail: Thumbnail{
					URL: IfNilPlaceholder(&d.Product.ImageUrl),
				},
				Author: Author{},
				Fields: []Fields{
					Fields{
						Name:   "Size",
						Value:  IfNilEmpty(&d.Product.Size),
						Inline: false,
					},
					Fields{
						Name:   "Payment",
						Value:  "||" + IfNilEmpty(&d.Product.Payment) + "||",
						Inline: false,
					},
					Fields{
						Name:   "Website",
						Value:  IfNilEmpty(&d.Product.Store),
						Inline: false,
					},
					Fields{
						Name:   "Mode",
						Value:  "Normal",
						Inline: false,
					},
					Fields{
						Name:   "Profile",
						Value:  "||" + IfNilEmpty(&d.Profile.ProfileName) + "||",
						Inline: false,
					},
					Fields{
						Name:   "Account",
						Value:  "||" + IfNilEmpty(&d.Profile.Email) + "||",
						Inline: false,
					},
				},
			},
		},
	}

	public := Webhook{
		Content:   "",
		Username:  "AresBot",
		AvatarURL: nil,
		Tts:       false,
		Embeds: []Embeds{
			Embeds{
				Title:       "Success",
				Description: "Checked out " + IfNilEmpty(&d.Product.Name),
				URL:         IfNilPlaceholder(&d.Product.Url),
				Timestamp:   &t,
				Color:       850000,
				Footer: Footer{
					Text:    "AresBot " + w.Version,
					IconURL: nil,
				},
				Image: Image{},
				Thumbnail: Thumbnail{
					URL: IfNilPlaceholder(&d.Product.ImageUrl),
				},
				Author: Author{},
				Fields: []Fields{
					Fields{
						Name:   "Size",
						Value:  IfNilEmpty(&d.Product.Size),
						Inline: false,
					},
					Fields{
						Name:   "Payment",
						Value:  IfNilEmpty(&d.Product.Payment),
						Inline: false,
					},
					Fields{
						Name:   "Website",
						Value:  IfNilEmpty(&d.Product.Store),
						Inline: false,
					},
					Fields{
						Name:   "Mode",
						Value:  "Normal",
						Inline: false,
					},
					Fields{
						Name:  "Quicktask",
						Value: "[Click](http://localhost:11914/quicktask?product=" + url.QueryEscape(IfNilEmpty(&d.Product.Url)) + "&size=" + IfNilEmpty(&d.Product.Size) + "&store=" + IfNilEmpty(&d.Product.Store) + ")",
					},
				},
			},
		},
	}

	checkoutlog := Webhook{
		Content:   "",
		Username:  "AresBot",
		AvatarURL: nil,
		Tts:       false,
		Embeds: []Embeds{
			Embeds{
				Title:       "User Checkout",
				Description: "Successful Checkout for " + IfNilEmpty(&d.Product.Name),
				URL:         IfNilPlaceholder(&d.Product.CheckoutUrl),
				Timestamp:   &t,
				Color:       850000,
				Footer: Footer{
					Text:    "AresBot " + IfNilEmpty(&w.Version),
					IconURL: nil,
				},
				Image: Image{},
				Thumbnail: Thumbnail{
					URL: IfNilPlaceholder(&d.Product.ImageUrl),
				},
				Author: Author{},
				Fields: []Fields{
					Fields{
						Name:   "Username",
						Value:  IfNilEmpty(&user),
						Inline: false,
					},
					Fields{
						Name:   "License Type",
						Value:  IfNilEmpty(&plan),
						Inline: false,
					},
					Fields{
						Name:   "License",
						Value:  IfNilEmpty(&w.HyperData.Key),
						Inline: false,
					},
					Fields{
						Name:   "Profile",
						Value:  "||" + IfNilEmpty(&d.Profile.ProfileName) + "||",
						Inline: false,
					},
					Fields{
						Name:   "Account",
						Value:  "||" + IfNilEmpty(&d.Profile.Email) + "||",
						Inline: false,
					},
				},
			},
		},
	}

	data_pub, err := json.Marshal(public)
	if err != nil {
		l.ErrorLogger.Println(err)
	}

	data_privat, err := json.Marshal(privat)
	if err != nil {
		l.ErrorLogger.Println(err)
	}

	data_checkout, err := json.Marshal(checkoutlog)
	if err != nil {
		l.ErrorLogger.Println(err)
	}

	w.PublicUrl = "https://discord.com/api/webhooks/1068236706404380812/H6gUhq67Yu5-quQya34nsK6CbtuFHccCrwMIEFH0fBj0d29A2B46gaaMBp6z4uOg4OCI"
	_, err = c.Post(w.PublicUrl, "application/json", bytes.NewBuffer(data_pub))
	if err != nil {
		l.ErrorLogger.Println(err)
	}

	// fallback
	privaturl := w.PrivatUrl
	if privaturl == "" {
		privaturl = loadFallbackWebhook()
	}

	_, err = c.Post(privaturl, "application/json", bytes.NewBuffer(data_privat))
	if err != nil {
		l.ErrorLogger.Println(err)
	}
	_, err = c.Post(w.PrivatUrl, "application/json", bytes.NewBuffer(data_checkout))
	if err != nil {
		l.ErrorLogger.Println(err)
	}
}

func (w *WebhookHandler) SendCustomWebhook(public Webhook, privat Webhook, checkoutlog Webhook) {
	c := http.Client{}

	data_pub, err := json.Marshal(public)
	if err != nil {
		l.ErrorLogger.Println(err)
	}

	data_privat, err := json.Marshal(privat)
	if err != nil {
		l.ErrorLogger.Println(err)
	}

	data_checkout, err := json.Marshal(checkoutlog)
	if err != nil {
		l.ErrorLogger.Println(err)
	}

	_, err = c.Post("https://discord.com/api/webhooks/1068236706404380812/H6gUhq67Yu5-quQya34nsK6CbtuFHccCrwMIEFH0fBj0d29A2B46gaaMBp6z4uOg4OCI", "application/json", bytes.NewBuffer(data_pub))
	if err != nil {
		l.ErrorLogger.Println(err)
	}

	// fallback
	privaturl := w.PrivatUrl
	if privaturl == "" {
		privaturl = loadFallbackWebhook()
	}

	_, err = c.Post(privaturl, "application/json", bytes.NewBuffer(data_privat))
	if err != nil {
		l.ErrorLogger.Println(err)
	}
	_, err = c.Post("https://discord.com/api/webhooks/936380184762933320/016In92YFNCARaU7MqfisWeP-V8K4ErRr3ZzgHasNiCpTIP7O4PiOQDXIocV2W2YD3hK", "application/json", bytes.NewBuffer(data_checkout))
	if err != nil {
		l.ErrorLogger.Println(err)
	}
}

func loadFallbackWebhook() string {
	var settings models.Settings
	file, err := ioutil.ReadFile("settings.json")
	if err != nil {
		l.ErrorLogger.Println(err)

		fmt.Println("Couldn't find settings.json, closing..")

		time.Sleep(5 * time.Second)
		os.Exit(-1)
	}
	err = json.Unmarshal(file, &settings)
	if err != nil {
		l.ErrorLogger.Println(err)

		fmt.Println("Error in file settings.json, closing.. ")

		time.Sleep(5 * time.Second)
		fmt.Println(err)
	}

	return settings.WebhookURL
}

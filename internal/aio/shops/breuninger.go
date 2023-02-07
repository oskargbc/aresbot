package shops

import (
	"aresbot/internal/aio/models"
	"aresbot/internal/aio/shops/breuninger"
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
)

func NewBreuningerBot() ShopBot {
	return &Breuninger{}
}

type Breuninger struct {
	ShopBot

	Client          *http.Client
	SizeVariant     string
	AtcUrl          string
	VertriebsinfoId string
	WaitCh          chan int
	Referer         string
}

const (
	atcUrI = "https://www.breuninger.com/de/kauf/addToCart"
)

func (m *Breuninger) CheckProduct(task models.Task, proxy *models.Proxy) Response {
	if proxy != nil {
		m.Client.Transport = &http.Transport{Proxy: http.ProxyURL(&proxy.Url)}
	}

	// get page
	req, err := http.NewRequest(http.MethodGet, task.ProductURL, nil)
	if err != nil {
		return Response{
			e:       nil,
			success: false,
		}
	}

	for k, v := range breuninger.LoadPageHeader {
		req.Header.Add(k, v)
	}

	resp, err := m.Client.Do(req)
	if err != nil {
		return Response{
			e:       err,
			success: false,
		}
	}

	switch resp.StatusCode {
	case 200:
		break
	case 403:
		return Response{
			e:           errors.New("Proxy blocked"),
			rotateProxy: true,
			success:     false,
		}
	default:
		return Response{
			e:           errors.New("Unexpected statuscode"),
			rotateProxy: true,
			success:     false,
		}
	}

	// parse variants
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return Response{
			e:       err,
			success: false,
		}
	}
	body, ok := doc.Find("body > main > div > bewerten-vue-data").Attr("data-bewerten-json-content")
	if !ok {
		return Response{
			e:       errors.New("didn't find json content"),
			success: false,
		}
	}

	brand := doc.Find("#bewerten-vue-pds > section > div.shop-grid-column.shop-grid-column--24.shop-grid-column--l-12.shop-grid-column--xl-12.shop-grid-column--xxl-9.hidden--L_and_up > div.bewerten-zusammenfassung > h1 > a > span").Text()
	vari := doc.Find("#bewerten-vue-pds > section > div.shop-grid-column.shop-grid-column--24.shop-grid-column--l-12.shop-grid-column--xl-12.shop-grid-column--xxl-9.hidden--L_and_up > div.bewerten-zusammenfassung > h1 > div > span").Text()

	var bProduct map[string]interface{}
	_ = json.Unmarshal([]byte(body), &bProduct)

	if len(bProduct) == 0 {
		return Response{
			e: errors.New("not found product content"),
			//rotateProxy: false,
			success: false,
		}
	}

	var BreuProduct breuninger.BreunigerArtikel

	uri, err := url.Parse(task.ProductURL)
	if err != nil {
		return Response{
			e: err,
			//rotateProxy: false,
			success: false,
		}
	}

	q := uri.Query()

	urisplit := ""
	if q.Has("variant") {
		urisplit = q.Get("variant")
	}

	for artikel, avalue := range bProduct["artikel"].(map[string]interface{}) {
		if urisplit == "" {
			b, _ := json.Marshal(avalue)
			err = json.Unmarshal(b, &BreuProduct)
			if err != nil {
				return Response{
					e: err,
					//rotateProxy: false,
					success: false,
				}
			}

			break
		} else {
			if artikel == urisplit {
				b, _ := json.Marshal(avalue)
				err = json.Unmarshal(b, &BreuProduct)
				if err != nil {
					return Response{
						e: err,
						//rotateProxy: false,
						success: false,
					}
				}

				break
			}
		}
	}

	sizes := BreuProduct.Groessen.Werte

	size := task.Size
	if task.Size == "random" || task.Size == "" {
		size = sizes[rand.Intn(len(sizes)-1)].Name
	}

	for _, n := range sizes {
		if size == n.Name {
			m.SizeVariant = n.ArtikelNr
		}
	}

	p := models.Product{
		Name:     brand + vari,
		Size:     size,
		Url:      task.ProductURL,
		ImageUrl: BreuProduct.Ansichten[0].ThumbnailUrl,
		Payment:  "",
		Store:    task.Store,
	}

	// request size url to get vertriebsinfoId:string

	atcUrl := "https://" + uri.Host + uri.Path + "?variant=" + m.SizeVariant
	m.Referer = atcUrl

	req, err = http.NewRequest(http.MethodGet, atcUrl, nil)

	if err != nil {
		return Response{
			e:       nil,
			success: false,
		}
	}

	for k, v := range breuninger.LoadPageHeader {
		req.Header.Add(k, v)
	}

	resp, err = m.Client.Do(req)
	if err != nil {
		return Response{
			e:       err,
			success: false,
		}
	}

	switch resp.StatusCode {
	case 200:
		break
	case 403:
		return Response{
			e:           errors.New("Proxy blocked"),
			rotateProxy: true,
			success:     false,
		}
	default:
		return Response{
			e:           errors.New("Unexpected statuscode"),
			rotateProxy: true,
			success:     false,
		}
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{
			e:       err,
			success: false,
		}
	}

	m.WaitCh = make(chan int, 1)
	getId := strings.Split(string(bodyBytes), "vertriebsinfoId&quot;:&quot;")
	if len(getId) == 0 {
		return Response{
			e:       errors.New("didn't find vertriebsinfoId"),
			success: false,
		}
	}
	m.VertriebsinfoId = strings.Split(getId[1], "&quot;,&quot;kreditorId&quot;")[0]
	m.Client.Jar.SetCookies(resp.Request.URL, resp.Cookies())

	return Response{
		e: nil,
		//rotateProxy: false,
		success: true,
		product: &p,
	}
}

func (m *Breuninger) GenerateCookies(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}
func (m *Breuninger) Login(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}
func (m *Breuninger) AddToCart(task models.Task, proxy *models.Proxy) Response {
	if proxy != nil {
		m.Client.Transport = &http.Transport{Proxy: http.ProxyURL(&proxy.Url)}
	}

	var data = strings.NewReader(`{"nummer":"` + m.SizeVariant + `","vertriebsinfoId":"` + m.VertriebsinfoId + `","anzahl":1}`)
	req, err := http.NewRequest(http.MethodPost, atcUrI, data)
	if err != nil {
		return Response{e: err, success: false}
	}
	for k, v := range breuninger.BreunigerAddToCart {
		req.Header.Add(k, v)
	}
	req.Header.Add("referer", m.Referer)

	resp, err := m.Client.Do(req)
	if err != nil {
		return Response{
			e:       err,
			success: false,
		}
	}
	b, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	var atcResponse breuninger.BreunigerAtcResponse
	err = json.Unmarshal(b, &atcResponse)
	if err != nil {
		return Response{
			e:       err,
			success: false,
		}
	}

	if strings.ToLower(atcResponse.Nachricht) == "ok" {
		m.Client.Jar.SetCookies(resp.Request.URL, resp.Cookies())
		return Response{
			//rotateProxy: false,
			success: true,
		}
	}

	return Response{
		e: errors.New("atc failed"),
		//rotateProxy: false,
		success: false,
	}

}

func (m *Breuninger) SubmitShipping(task models.Task, proxy *models.Proxy) Response {
	if proxy != nil {
		m.Client.Transport = &http.Transport{Proxy: http.ProxyURL(&proxy.Url)}
	}

	//get csrf
	req, err := http.NewRequest(http.MethodGet, "https://www.breuninger.com/de/kauf/contact", nil)
	if err != nil {
		return Response{e: err, success: false}
	}
	for k, v := range breuninger.BreunigerAddToCart {
		req.Header.Add(k, v)
	}

	resp, err := m.Client.Do(req)
	if err != nil {
		return Response{
			e:       err,
			success: false,
		}
	}

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	csrf, ok := doc.Find("#_csrf").Attr("value")
	if !ok {
		return Response{
			e: errors.New("didn't found csrf"),
			//rotateProxy: false,
			success: false,
		}
	}
	m.Client.Jar.SetCookies(resp.Request.URL, resp.Cookies())

	date := "0" + strconv.Itoa(rand.Intn(9)) + ".0" + strconv.Itoa(rand.Intn(9)) + ".1999"

	var data = strings.NewReader(`anrede=Herr&titel=&vorname=` + url.QueryEscape(task.Profile.FirstName) + `&nachname=` + url.QueryEscape(task.Profile.LastName) + `&strasse=` + url.QueryEscape(task.Profile.Address) + `&hausnummer=` + task.Profile.Address2 + `&adresszusatz=&plz=` + task.Profile.Zip + `&ort=` + url.QueryEscape(task.Profile.City) + `&landiso3=` + task.Profile.Country + `&email=` + task.Profile.Email + `&telefon=` + task.Profile.Phone + `&geburtsdatum=` + date + `&_csrf=` + csrf + `&marktforschung=off`)

	req, err = http.NewRequest("POST", "https://www.breuninger.com/de/kauf/contact", data)
	if err != nil {
		return Response{e: err}
	}

	for k, v := range breuninger.BreuningerShippingHeader {
		req.Header.Add(k, v)
	}

	resp, err = m.Client.Do(req)
	if err != nil {
		return Response{e: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.Request.URL.String() != "https://www.breuninger.com/de/kauf/delivery" {
		return Response{
			e:       errors.New("unexpected statuscode, " + strconv.Itoa(resp.StatusCode)),
			success: false,
		}
	}
	m.Client.Jar.SetCookies(resp.Request.URL, resp.Cookies())

	data = strings.NewReader(`versandOption=STANDARD_HERMES&lieferorttyp=`)
	req, err = http.NewRequest("POST", "https://www.breuninger.com/de/kauf/delivery", data)
	if err != nil {
		return Response{e: err}
	}

	for k, v := range breuninger.BreuningerShippingHeader2 {
		req.Header.Add(k, v)
	}

	resp, err = m.Client.Do(req)
	if err != nil {
		return Response{e: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 && resp.Request.URL.String() == "https://www.breuninger.com/de/kauf/paymentMethod" {
		m.Client.Jar.SetCookies(resp.Request.URL, resp.Cookies())
		return Response{
			//rotateProxy: false,
			success: true,
		}
	}
	return Response{
		e:       err,
		success: false,
	}

}
func (m *Breuninger) SubmitBilling(task models.Task, proxy *models.Proxy) Response {
	if proxy != nil {
		m.Client.Transport = &http.Transport{Proxy: http.ProxyURL(&proxy.Url)}
	}

	var data = strings.NewReader(`zahlart=paypal&kartenNummer=&pin=`)
	req, err := http.NewRequest("POST", "https://www.breuninger.com/de/kauf/paymentMethod", data)
	if err != nil {
		return Response{e: err}
	}

	for k, v := range breuninger.BreuningerBillingHeader {
		req.Header.Add(k, v)
	}

	resp, err := m.Client.Do(req)
	if err != nil {
		return Response{e: err}
	}
	defer resp.Body.Close()

	m.Client.Jar.SetCookies(resp.Request.URL, resp.Cookies())

	return Response{
		e: nil,
		//rotateProxy: false,
		success: true,
	}
}

func (m *Breuninger) Checkout(task models.Task, proxy *models.Proxy, product models.Product) Response {
	if proxy != nil {
		m.Client.Transport = &http.Transport{Proxy: http.ProxyURL(&proxy.Url)}
	}

	var data = strings.NewReader(`lieferzeitenJson=`)
	req, err := http.NewRequest("POST", "https://www.breuninger.com/de/kauf/paypal", data)
	if err != nil {
		return Response{e: err}
	}

	for k, v := range breuninger.BreuningerCoPaypalHeader {
		req.Header.Add(k, v)
	}

	resp, err := m.Client.Do(req)
	if err != nil {
		return Response{e: err}
	}
	defer resp.Body.Close()

	product.CheckoutUrl = resp.Request.URL.String()
	product.Payment = "PAYPAL"

	return Response{
		e:       nil,
		success: true,
		product: &product,
	}
}

func (m *Breuninger) NeedLogin(task models.Task) bool {
	return false
}
func (m *Breuninger) NeedCookies(task models.Task) bool {
	return false
}
func (m *Breuninger) GetSettings(settings models.Settings) error {
	j, _ := cookiejar.New(nil)

	m.Client = &http.Client{
		Jar: j,
	}

	return nil
}
func (m *Breuninger) IsActive() bool {
	return true
}

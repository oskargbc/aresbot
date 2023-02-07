package footasylum

func NewCheckProductHeader(Url string) map[string]string {

	return map[string]string{
		"sec-ch-ua":          "Not A;Brand;v=99, Chromium;v=96, Google Chrome;v=96",
		"sec-ch-ua-mobile":   "?0",
		"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36",
		"Content-Type":       "application/json",
		"Accept":             "*/*",
		"Referer":            Url,
		"X-Requested-With":   "XMLHttpRequest",
		"sec-ch-ua-platform": "macOS",
	}
}

func NewAddToCart(Url string) map[string]string {

	return map[string]string{
		"authority":          "www.footpatrol.de",
		"sec-ch-ua":          "Not A;Brand;v=99, Chromium;v=96, Google Chrome;v=96",
		"accept":             "*/*",
		"content-type":       "application/json",
		"x-requested-with":   "XMLHttpRequest",
		"sec-ch-ua-mobile":   "?0",
		"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36",
		"sec-ch-ua-platform": "macOS",
		"origin":             "https://www.footpatrol.de",
		"sec-fetch-site":     "same-origin",
		"sec-fetch-mode":     "cors",
		"sec-fetch-dest":     "empty",
		"referer":            Url,
	}
}

var CheckoutPaypalHeaders = map[string]string{
	"authority":          "www.footpatrol.de",
	"content-length":     "0",
	"sec-ch-ua":          "Not A;Brand;v=99, Chromium;v=96, Google Chrome;v=96",
	"accept":             "application/json",
	"sec-ch-ua-mobile":   "?0",
	"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36",
	"sec-ch-ua-platform": "macOS",
	"origin":             "https://www.footpatrol.de",
	"sec-fetch-site":     "same-origin",
	"sec-fetch-mode":     "cors",
	"sec-fetch-dest":     "empty",
	"referer":            "https://www.footpatrol.de/cart/",
	"accept-language":    "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
}

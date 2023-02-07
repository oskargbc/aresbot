package otto

var OttoLoginHeaders = map[string]string{
	"content-type":    "application/x-www-form-urlencoded",
	"accept-encoding": "gzip, deflate, br",
	"origin":          "https://www.otto.de",
	"referer":         "https://www.otto.de/user/login",
	"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"user-agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36",
	"accept-language": "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
}

var OttoLoadProductHeaders = map[string]string{
	"accept-encoding": "gzip, deflate, br",
	"referer":         "https://www.otto.de/user/login",
	"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"user-agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36",
	"accept-language": "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
}

var OttoAtcHeaders = map[string]string{
	"authority":          "www.otto.de",
	"accept":             "*/*",
	"content-type":       "application/json; charset=UTF-8",
	"accept-encoding":    "gzip, deflate, br",
	"origin":             "https://www.otto.de",
	"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36",
	"accept-language":    "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	"sec-ch-ua":          "\"Google Chrome\";v=\"95\", \"Chromium\";v=\"95\", \";Not A Brand\";v=\"99\"",
	"x-requested-with":   "XMLHttpRequest",
	"sec-ch-ua-mobile":   "?0",
	"sec-ch-ua-platform": "\"macOS\"",
	"sec-fetch-site":     "same-origin",
	"sec-fetch-mode":     "cors",
	"sec-fetch-dest":     "empty",
}

var OttoPaymentHeaders = map[string]string{
	"content-type":    "application/x-www-form-urlencoded",
	"accept-encoding": "gzip, deflate, br",
	"origin":          "https://www.otto.de",
	"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"user-agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36",
	"accept-language": "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	"referer":         "https://www.otto.de/order/checkout",
}

var OttoCheckoutOrderHeaderInit = map[string]string{
	"referer":         "https://www.otto.de/user/login",
	"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"user-agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36",
	"accept-language": "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
}

var OttoCheckoutOrderHeader = map[string]string{
	"authority":          "www.otto.de",
	"accept":             "*/*",
	"content-type":       "application/x-www-form-urlencoded; charset=UTF-8",
	"accept-encoding":    "gzip, deflate, br",
	"origin":             "https://www.otto.de",
	"referer":            "https://www.otto.de/order/checkout",
	"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36",
	"accept-language":    "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	"sec-ch-ua":          "\"Google Chrome\";v=\"95\", \"Chromium\";v=\"95\", \";Not A Brand\";v=\"99\"",
	"x-requested-with":   "XMLHttpRequest",
	"sec-ch-ua-mobile":   "?0",
	"sec-ch-ua-platform": "\"macOS\"",
	"sec-fetch-site":     "same-origin",
	"sec-fetch-mode":     "cors",
	"sec-fetch-dest":     "empty",
}

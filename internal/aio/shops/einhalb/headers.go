package einhalb

var GetProductHeader = map[string]string{
	"authority":                 "www.43einhalb.com",
	"cache-control":             "max-age=0",
	"sec-ch-ua":                 "Not;A Brand;v=99, Google Chrome;v=97, Chromium;v=97",
	"sec-ch-ua-mobile":          "?0",
	"sec-ch-ua-platform":        "macOS",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36",
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"sec-fetch-site":            "none",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-user":            "?1",
	"sec-fetch-dest":            "document",
	"accept-language":           "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
}

var ProductAddToCartHeader = map[string]string{
	"authority":          "www.43einhalb.com",
	"sec-ch-ua":          "Not;A Brand;v=99, Google Chrome;v=97, Chromium;v=97",
	"accept":             "*/*",
	"content-type":       "application/x-www-form-urlencoded; charset=UTF-8",
	"x-requested-with":   "XMLHttpRequest",
	"sec-ch-ua-mobile":   "?0",
	"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36",
	"sec-ch-ua-platform": "macOS",
	"origin":             "https://www.43einhalb.com",
	"sec-fetch-site":     "same-origin",
	"sec-fetch-mode":     "cors",
	"sec-fetch-dest":     "empty",
	"accept-language":    "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
}

var PaypalCreatePaymentHeaders = map[string]string{
	"authority":          "www.43einhalb.com",
	"content-length":     "0",
	"sec-ch-ua":          "Not;A Brand;v=99, Google Chrome;v=97, Chromium;v=97",
	"accept":             "application/json",
	"sec-ch-ua-mobile":   "?0",
	"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36",
	"sec-ch-ua-platform": "macOS",
	"origin":             "https://www.43einhalb.com",
	"sec-fetch-site":     "same-origin",
	"sec-fetch-mode":     "cors",
	"sec-fetch-dest":     "empty",
	"referer":            "https://www.43einhalb.com/warenkorb",
	"accept-language":    "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
}
var PaypalCheckoutHeaders = map[string]string{
	"authority":          "www.paypal.com",
	"sec-ch-ua":          "Not;A Brand;v=99, Google Chrome;v=97, Chromium;v=97",
	"sec-ch-ua-mobile":   "?0",
	"x-requested-by":     "smart-payment-buttons",
	"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36",
	"content-type":       "application/json",
	"accept":             "application/json",
	"x-csrf-jwt":         "__blank__",
	"x-requested-with":   "XMLHttpRequest",
	"x-cookies":          "{}",
	"sec-ch-ua-platform": "macOS",
	"origin":             "https://www.paypal.com",
	"sec-fetch-site":     "same-origin",
	"sec-fetch-mode":     "cors",
	"sec-fetch-dest":     "empty",
	"referer":            "https://www.paypal.com/smart/button?env=production&logLevel=warn&style.color=blue&style.shape=rect&style.size=responsive&style.tagline=false&style.label=checkout&locale.x=de_DE&domain=www.43einhalb.com&sessionID=uid_f42d0af53e_mtq6mjc6mza&buttonSessionID=uid_a7ceb8ec80_mtq6mjc6mza&renderedButtons=paypal&storageID=uid_45bb1cf24b_mja6mze6nde&funding.disallowed=venmo&funding.remembered=paypal&sdkMeta=eyJ1cmwiOiJodHRwczovL3d3dy5wYXlwYWxvYmplY3RzLmNvbS9hcGkvY2hlY2tvdXQuanMifQ&uid=6ceccf2e4d&version=4&xcomponent=1",
	"accept-language":    "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
}

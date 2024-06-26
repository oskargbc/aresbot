package breuninger

var LoadLoginCsrfHeader = map[string]string{
	"authority":          "www.breuninger.com",
	"sec-ch-ua":          "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"99\", \"Google Chrome\";v=\"99\"",
	"sec-ch-ua-mobile":   "?0",
	"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36",
	"sec-ch-ua-platform": "\"macOS\"",
	"accept":             "*/*",
	"sec-fetch-site":     "same-origin",
	"sec-fetch-mode":     "cors",
	"sec-fetch-dest":     "empty",
	"referer":            "https://www.breuninger.com/de/kauf/login/konto?referrer=/de/account/overview",
	"accept-language":    "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
}

var LoadPageHeader = map[string]string{
	"authority":                 "www.breuninger.com",
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"accept-language":           "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	"cache-control":             "max-age=0",
	"if-none-match":             "'myra-beab399c'",
	"sec-ch-ua":                 "'Not_A Brand';v='99', 'Google Chrome';v='109', 'Chromium';v='109'",
	"sec-ch-ua-mobile":          "?0",
	"sec-ch-ua-platform":        "'macO'",
	"sec-fetch-dest":            "document",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-site":            "none",
	"sec-fetch-user":            "?1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
}

var BreunigerAddToCart = map[string]string{
	"authority":          "www.breuninger.com",
	"accept":             "*/*",
	"accept-language":    "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	"content-type":       "application/vnd.position+json",
	"origin":             "https://www.breuninger.com",
	"sec-ch-ua":          "'Not_A Brand';v='99', 'Google Chrome';v='109', 'Chromium';v='109'",
	"sec-ch-ua-mobile":   "?0",
	"sec-ch-ua-platform": "'macOS'",
	"sec-fetch-dest":     "empty",
	"sec-fetch-mode":     "cors",
	"sec-fetch-site":     "same-origin",
	"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	"x-requested-with":   "XMLHttpRequest",
}

var BreuningerShippingHeader = map[string]string{
	"authority":                 "www.breuninger.com",
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"accept-language":           "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	"cache-control":             "max-age=0",
	"content-type":              "application/x-www-form-urlencoded",
	"origin":                    "https://www.breuninger.com",
	"referer":                   "https://www.breuninger.com/de/kauf/contact",
	"sec-ch-ua":                 `"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"`,
	"sec-ch-ua-mobile":          "?0",
	"sec-ch-ua-platform":        `"macOS"`,
	"sec-fetch-dest":            "document",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-site":            "same-origin",
	"sec-fetch-user":            "?1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7, AppleWebKit/537.36 (KHTML, like Gecko, Chrome/109.0.0.0 Safari/537.36",
}

var BreuningerShippingHeader2 = map[string]string{
	"authority":                 "www.breuninger.com",
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"accept-language":           "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	"cache-control":             "max-age=0",
	"content-type":              "application/x-www-form-urlencoded",
	"origin":                    "https://www.breuninger.com",
	"referer":                   "https://www.breuninger.com/de/kauf/delivery",
	"sec-ch-ua":                 `"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"`,
	"sec-ch-ua-mobile":          "?0",
	"sec-ch-ua-platform":        `"macOS"`,
	"sec-fetch-dest":            "document",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-site":            "same-origin",
	"sec-fetch-user":            "?1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7, AppleWebKit/537.36 (KHTML, like Gecko, Chrome/109.0.0.0 Safari/537.36",
}

var BreuningerBillingHeader = map[string]string{
	"authority":                 "www.breuninger.com",
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"accept-language":           "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	"cache-control":             "max-age=0",
	"content-type":              "application/x-www-form-urlencoded",
	"origin":                    "https://www.breuninger.com",
	"referer":                   "https://www.breuninger.com/de/kauf/paymentMethod",
	"sec-ch-ua":                 `"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"`,
	"sec-ch-ua-mobile":          "?0",
	"sec-ch-ua-platform":        `"macOS"`,
	"sec-fetch-dest":            "document",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-site":            "same-origin",
	"sec-fetch-user":            "?1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7, AppleWebKit/537.36 (KHTML, like Gecko, Chrome/109.0.0.0 Safari/537.36",
}

var BreuningerCoPaypalHeader = map[string]string{
	"authority":                 "www.breuninger.com",
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"accept-language":           "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	"cache-control":             "max-age=0",
	"content-type":              "application/x-www-form-urlencoded",
	"origin":                    "https://www.breuninger.com",
	"referer":                   "https://www.breuninger.com/de/kauf/lastCheck",
	"sec-ch-ua":                 `"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"`,
	"sec-ch-ua-mobile":          "?0",
	"sec-ch-ua-platform":        `"macOS"`,
	"sec-fetch-dest":            "document",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-site":            "same-origin",
	"sec-fetch-user":            "?1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7, AppleWebKit/537.36 (KHTML, like Gecko, Chrome/109.0.0.0 Safari/537.36",
}

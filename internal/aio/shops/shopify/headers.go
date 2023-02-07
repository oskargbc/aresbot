package shopify

import (
	"aresbot/internal/aio/shared"
	"fmt"
)

var PageLoadHeader = map[string]interface{}{
	"pragma":                    "no-cache",
	"cache-control":             "no-cache",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36",
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"sec-fetch-site":            "none",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-user":            "?1",
	"sec-fetch-dest":            "document",
	"accept-language":           "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
}

func SelectAddressHeaders(url string) map[string]string {
	hostName := shared.GetHostNameFromUrl(url)
	hostWithProtocol := fmt.Sprintf("https://%s", hostName)

	return map[string]string{
		"authority":                 hostName,
		"cache-control":             "max-age=0",
		"upgrade-insecure-requests": "1",
		"origin":                    hostWithProtocol,
		"content-type":              "application/x-www-form-urlencoded",
		"user-agent":                "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"sec-fetch-site":            "same-origin",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-user":            "?1",
		"sec-fetch-dest":            "document",
		"referer":                   hostWithProtocol,
		"accept-language":           "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	}
}

func SelectShippingHeaders(url string) map[string]string {
	hostName := shared.GetHostNameFromUrl(url)
	hostWithProtocol := fmt.Sprintf("https://%s", hostName)

	return map[string]string{
		"authority":                 hostName,
		"cache-control":             "max-age=0",
		"upgrade-insecure-requests": "1",
		"origin":                    hostWithProtocol,
		"content-type":              "application/x-www-form-urlencoded",
		"user-agent":                "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"sec-fetch-site":            "same-origin",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-user":            "?1",
		"sec-fetch-dest":            "document",
		"referer":                   hostWithProtocol,
		"accept-language":           "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	}
}

func CheckoutHeaders(url string) map[string]string {
	hostName := shared.GetHostNameFromUrl(url)
	hostWithProtocol := fmt.Sprintf("https://%s", hostName)

	return map[string]string{
		"authority":                 hostName,
		"cache-control":             "max-age=0",
		"upgrade-insecure-requests": "1",
		"origin":                    hostWithProtocol,
		"content-type":              "application/x-www-form-urlencoded",
		"user-agent":                "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"sec-fetch-site":            "same-origin",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-user":            "?1",
		"sec-fetch-dest":            "document",
		"referer":                   hostWithProtocol,
		"accept-language":           "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	}
}

func PaymentHeaders(url string) map[string]string {
	hostName := shared.GetHostNameFromUrl(url)
	hostWithProtocol := fmt.Sprintf("https://%s", hostName)

	return map[string]string{
		"authority":                 hostName,
		"cache-control":             "max-age=0",
		"upgrade-insecure-requests": "1",
		"origin":                    hostWithProtocol,
		"content-type":              "application/x-www-form-urlencoded",
		"user-agent":                "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"sec-fetch-site":            "same-origin",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-user":            "?1",
		"sec-fetch-dest":            "document",
		"referer":                   hostWithProtocol,
		"accept-language":           "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	}
}

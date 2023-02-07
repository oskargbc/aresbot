package footlocker

import (
	"fmt"
)

func NewGenerateSessionHeaders(Site string, region string, u string) map[string]string {
	return map[string]string{
		"Accept":          "application/json",
		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.128 Safari/537.36",
		"X-Fl-Request-Id": u,
		"Origin":          fmt.Sprintf("https://www.%s.%s", Site, region),
		"Sec-Fetch-Site":  "same-origin",
		"Sec-Fetch-Mode":  "cors",
		"Sec-Fetch-Dest":  "empty",
		"Referer":         fmt.Sprintf("https://www.%s.%s", Site, region),
		"Accept-Language": "zh-CN,zh;q=0.9",
	}
}

func NewGetSizeHeaders(Site string, region string, u string, session string) map[string]string {
	return map[string]string{
		"Accept":             "application/json",
		"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.128 Safari/537.36",
		"X-Fl-Request-Id":    u,
		"X-Flapi-session.id": session,
		"Origin":             fmt.Sprintf("https://www.%s.%s", Site, region),
		"Sec-Fetch-Site":     "same-origin",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Dest":     "empty",
		"Referer":            fmt.Sprintf("https://www.%s.%s", Site, region),
		"Accept-Language":    "zh-CN,zh;q=0.9",
		"Cookie":             "JSESSIONID=" + session,
	}
}
func NewAddToCartHeaders(CSRF string, SizeID string, JSESSIONID string, UUID string) map[string]string {
	return map[string]string{
		"authority":          "www.footlocker.de",
		"sec-ch-ua":          "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"",
		"x-csrf-token":       CSRF,
		"x-api-lang":         "de-DE",
		"accept-language":    "de-DE,de;q=0.8",
		"sec-ch-ua-mobile":   "?0",
		"x-fl-productid":     SizeID,
		"content-type":       "application/json",
		"accept":             "application/json",
		"x-flapi-session-id": JSESSIONID,
		"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36",
		"x-fl-request-id":    UUID,
		"sec-ch-ua-platform": "\"macOS\"",
		"origin":             "https://www.footlocker.de",
		"sec-fetch-site":     "same-origin",
		"sec-fetch-mode":     "cors",
		"sec-fetch-dest":     "empty",
		"cookie":             "JSESSIONID=" + JSESSIONID,
	}
}

var CheckProductHeaders = map[string]string{
	"User-Agent":    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36",
	"Content-Type":  "application/x-www-form-urlencoded; charset=UTF-8",
	"cache-control": "no-store,no-cache,must-revalidate,proxy-revalidate,max-age=0",
	"pragma":        "no-cache",
}

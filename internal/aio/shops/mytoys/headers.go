package mytoys

var MytoysLoadProductHeaders = map[string]string{
	"host":            "mspublicapi.apps.aws.mytoys.com",
	"user-agent":      "myToys/7.6.0 iOS/14.2.0 (iPhone10,5) CFNetwork/1206",
	"content-type":    "application/x-www-form-urlencoded",
	"accept":          "application/vnd.mytoys.app.product-1.0.0+json",
	"authorization":   "Basic aW9zOk42cmxvMGw1Qi4=",
	"accept-language": "de-de",
	"shopName":        "mytoys",
}

var MytoysLoginHeaders = map[string]string{
	"host":            "checkout.mytoys.de",
	"content-type":    "application/x-www-form-urlencoded",
	"origin":          "https://checkout.mytoys.de",
	"accept-encoding": "gzip, deflate, br",
	"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"user-agent":      "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
	"referer":         "https://checkout.mytoys.de/checkout/login",
	"accept-language": "de-de",
}

var MytoysAtcHeaders = map[string]string{
	"user-agent":      "myToys/9.2.0 iOS/14.4.0 (iPhone10,5) CFNetwork/1220.1",
	"accept":          "application/vnd.mytoys.app.user-1.0.0+json",
	"authorization":   "Basic aW9zOk42cmxvMGw1Qi4=",
	"accept-language": "de-de",
	"accept-encoding": "gzip, deflate, br",
	"shopName":        "mytoys",
	"content-type":    "application/json",
}

var MytoysSessionHeaders = map[string]string{
	"user-agent":      "myToys/9.2.0 iOS/14.4.0 (iPhone10,5) CFNetwork/1220.1",
	"accept":          "application/vnd.mytoys.app.user-1.0.0+json",
	"authorization":   "Basic aW9zOk42cmxvMGw1Qi4=",
	"accept-language": "de-de",
	"shopName":        "mytoys",
	"content-type":    "application/x-www-form-urlencoded",
	"accept-encoding": "gzip, deflate, br",
}

var MytoysApiOrderHeaders = map[string]string{
	"user-agent":      "Mozilla/5.0 (iPhone; CPU iPhone OS 14_7_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
	"accept":          "application/json, text/javascript, */*; q=0.01",
	"accept-language": "de-de",
	"accept-encoding": "gzip, deflate, br",
	"content-type":    "application/json",
	"origin":          "https: //checkout.mytoys.de",
	"host":            "checkout.mytoys.de",
	"referer":         "https://checkout.mytoys.de/checkout/orderOverview",
}

var MytoysOrderOverviewHeaders = map[string]string{
	"user-agent":      "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
	"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"accept-language": "de-de",
	"accept-encoding": "gzip, deflate, br",
	"host":            "checkout.mytoys.de",
}
var MytoysBasketClearItemHeaders = map[string]string{
	"user-agent":      "Mozilla/5.0 (iPhone; CPU iPhone OS 14_7_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
	"accept":          "application/json, text/plain, */*",
	"accept-language": "de-de",
	"accept-encoding": "gzip, deflate, br",
	"referer":         "https://checkout.mytoys.de/checkout/basketOverview",
	"host":            "checkout.mytoys.de",
}

var MytoysBasketOverviewHeaders = map[string]string{
	"user-agent":      "Mozilla/5.0 (iPhone; CPU iPhone OS 14_7_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
	"accept":          "application/json, text/plain, */*",
	"accept-language": "de-de",
	"accept-encoding": "gzip, deflate, br",
	"referer":         "https://checkout.mytoys.de/checkout/basketOverview",
	"host":            "checkout.mytoys.de",
}

var MytoysBasketInfoHeaders = map[string]string{
	"user-agent":      "myToys/10.1.1 iOS/14.7.1 (iPhone10,5) CFNetwork/1240.0.4",
	"accept-language": "de-de",
	"accept-encoding": "gzip, deflate, br",
	"accept":          "application/vnd.mytoys.app.user-1.0.0+json",
	"shopname":        "mytoys",
	"authorization":   "Basic aW9zOk42cmxvMGw1Qi4=",
	"content-type":    "application/x-www-form-urlencoded",
}

var MytoysPayOrderHeaders = map[string]string{
	"user-agent":      "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
	"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"accept-language": "de-de",
	"accept-encoding": "gzip, deflate, br",
	"host":            "checkout.mytoys.de",
	"referer":         "https://checkout.mytoys.de/checkout/orderOverview",
}

var MytoysCreditCardHeaders = map[string]string{
	"user-agent":      "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
	"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"accept-language": "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	"accept-encoding": "gzip, deflate, br",
	"content-type":    "application/x-www-form-urlencoded",
	"host":            "www.computop-paygate.com",
	"origin":          "https://www.computop-paygate.com",
}

var MytoysThreeDsHeaders = map[string]interface{}{
	"user-agent":      "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
	"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"accept-language": "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	"accept-encoding": "gzip, deflate, br",
	"origin":          "https://www.computop-paygate.com",
}

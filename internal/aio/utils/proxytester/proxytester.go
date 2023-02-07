package proxytester

import (
	"aresbot/internal/aio/models"
	"net/http"
	"time"
)

type ProxyTester struct {
	Proxys        []models.Proxy
	Url           string
	Timeout       int
	WorkingProxys []models.Proxy
}

func (p *ProxyTester) TestProxys() {
	client := http.Client{
		Timeout: time.Duration(p.Timeout) * time.Second,
	}

	for _, proxy := range p.Proxys {
		client.Transport = &http.Transport{Proxy: http.ProxyURL(&proxy.Url)}
	}
}

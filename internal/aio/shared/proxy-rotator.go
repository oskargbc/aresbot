package shared

import (
	"aresbot/internal/aio/models"
)

type ProxyRotator struct {
	Proxys []models.Proxy
}

func (p *ProxyRotator) RotateProxy() *models.Proxy {
	if len(p.Proxys) > 1 {
		prox := p.Proxys[0]

		p.Proxys = p.Proxys[1:]

		return &prox
	} else {
		return nil
	}
}

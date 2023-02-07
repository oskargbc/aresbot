package shops

import "aresbot/internal/aio/models"

func NewAirnessBot() ShopBot {
	return &Airness{}
}

type Airness struct {
	ShopBot
}

func (s *Airness) CheckProduct(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}
func (s *Airness) GenerateCookies(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}
func (s *Airness) Login(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}
func (s *Airness) AddToCart(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}

func (s *Airness) SubmitBilling(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}

func (s *Airness) SubmitShipping(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}

func (s *Airness) Checkout(task models.Task, proxy *models.Proxy, product models.Product) Response {
	return Response{}
}
func (s *Airness) NeedLogin(task models.Task) bool {
	return false
}
func (s *Airness) NeedCookies(task models.Task) bool {
	return false
}
func (s *Airness) GetSettings(settings models.Settings) error {
	return nil
}
func (s *Airness) IsActive() bool {
	return true
}

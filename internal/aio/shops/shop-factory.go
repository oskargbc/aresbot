package shops

import (
	models2 "aresbot/internal/aio/models"
	"log"
)

type ShopBot interface {
	CheckProduct(task models2.Task, proxy *models2.Proxy) Response
	GenerateCookies(task models2.Task, proxy *models2.Proxy) Response
	Login(task models2.Task, proxy *models2.Proxy) Response
	AddToCart(task models2.Task, proxy *models2.Proxy) Response
	SubmitBilling(task models2.Task, proxy *models2.Proxy) Response
	SubmitShipping(task models2.Task, proxy *models2.Proxy) Response
	Checkout(task models2.Task, proxy *models2.Proxy, product models2.Product) Response
	NeedLogin(task models2.Task) bool
	NeedCookies(task models2.Task) bool
	GetSettings(settings models2.Settings) error
	IsActive() bool
	log() *log.Logger
}

var ShopBots = map[string]ShopBot{
	"VOOSTORE":   NewVoostoreBot(),
	"BREUNINGER": NewBreuningerBot(),
}

/*

func NewMytoysBot() ShopBot {
	return &Mytoys{}
}

type Mytoys struct {
	ShopBot
}

func (m *Mytoys) CheckProduct(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}
func (m *Mytoys) GenerateCookies(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}
func (m *Mytoys) Login(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}
func (m *Mytoys) AddToCart(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}

func (m *Mytoys) SubmitBilling(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}

func (m *Mytoys) SubmitShipping(task models.Task, proxy *models.Proxy) Response {
	return Response{}
}

func (m *Mytoys) Checkout(task models.Task, proxy *models.Proxy, product models.Product) Response {
	return Response{}
}
func (m *Mytoys) NeedLogin(task models.Task) bool {
	return false
}
func (m *Mytoys) NeedCookies(task models.Task) bool {
	return false
}
func (m *Mytoys) GetSettings(settings models.Settings) error {
	return nil
}
func (m *Mytoys) IsActive() bool {
	return true
}

*/

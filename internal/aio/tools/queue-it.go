package tools

import (
	"aresbot/internal/aio/constants"
	"aresbot/internal/aio/models"
	"aresbot/internal/aio/shared"
	"aresbot/internal/aio/tools/queueit"
	"aresbot/pkg/logger"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync"
	"time"
)

func NewQueueItTool() Tool {
	return &QueueItTool{}
}

type QueueItTool struct {
	Tool
	Solver  shared.CaptchaSolver
	Handler shared.WebhookHandler
}

func (s *QueueItTool) Run(rotator shared.ProxyRotator, extras map[string]interface{}, handler shared.WebhookHandler) {
	entry := extras["entrys"]
	url := extras["queueUrl"]
	useProxy := extras["useProxy"]
	s.Handler = handler

	shop, c, e, t, err, a := parseUrlValues(*url.(*string))
	if err != nil {
		logger.ErrorLogger.Println(err)
		shared.OError(1, "Wrong queue url")
		return
	}

	wg := &sync.WaitGroup{}

	for i := 0; i < entry.(int); i++ {
		wg.Add(1)
		var proxy *models.Proxy

		if !useProxy.(bool) {
			proxy = rotator.RotateProxy()
		}

		go s.QueueItRunner(wg, proxy, shop, c, e, t, i+1, a)
	}

	wg.Wait()
	return
}

func parseUrlValues(s string) (string, string, string, string, error, string) {
	u, err := url.Parse(s)
	if err != nil {
		logger.WarningLogger.Println("couldn't parse queue url ", s)
		return "", "", "", "", err, s
	}
	shop := strings.Split(u.Host, ".")[0]

	q, _ := url.ParseQuery(u.RawQuery)
	if err != nil {
		logger.WarningLogger.Println("couldn't parse queue url querys ", s)
		return "", "", "", "", err, s
	}

	return shop, q["c"][0], q["e"][0], q["t"][0], nil, s
}

func (s *QueueItTool) IsActive() bool {
	return true
}

func (s *QueueItTool) GetSettings(settings models.Settings, id int) error {
	if len(settings.CaptchaKey) < 6 {
		return constants.NewGeneralError("no captcha key given")
	}

	p := "2captcha;anticaptcha"
	if !strings.Contains(p, settings.CaptchaProvider) {
		return constants.NewGeneralError("only 2captcha for anticaptcha given: " + settings.CaptchaProvider)
	}

	s.Solver = shared.NewCaptchaSolver(settings.CaptchaKey, settings.CaptchaProvider)

	return nil
}

func (s *QueueItTool) NeedBackend() bool {
	return true
}

func (s *QueueItTool) QueueItRunner(wg *sync.WaitGroup, proxy *models.Proxy, shop string, c string, e string, t string, id int, baseUrl string) {
	found := false
	params := queueit.Params{}

	shared.OInfo(id, "Waiting for backend data")

	for !found {
		resp, err := http.Get(constants.ApiToolBaseUrl + "queueit/params")
		if err != nil {
			logger.ErrorLogger.Println("coudn't reach api ", err)
		}

		body, _ := ioutil.ReadAll(resp.Body)

		var response queueit.ApiReponse
		_ = json.Unmarshal(body, &response)

		for _, v := range response.Drop {
			if strings.ToLower(shop) == strings.ToLower(v.Shop) {
				params = v.Params
			}
		}

		if params.LayoutName != "" {
			found = true
		} else {
			time.Sleep(time.Second * 2)
		}
	}

	shared.OInfo(id, "Starting.. ")

	target, _ := url.QueryUnescape(t)

	j, _ := cookiejar.New(nil)
	client := &http.Client{Jar: j}
	if proxy != nil {
		client.Transport = &http.Transport{Proxy: http.ProxyURL(&proxy.Url)}
	}

	manager := queueit.NewQueueItManager(queueit.GenerateTimestamp(), c, e, shop, params.Sitekey, target, params.LayoutName, params.LayoutVersion, client, s.Solver, id, s.Handler, baseUrl)

	shared.OInfo(id, "Generating cookie.. ")
	manager.GetEnterChallengeCookie(t)
	w := sync.WaitGroup{}

	go manager.GetRecaptchaChallange(&w)
	err := manager.GetPowChallenge()

	if err != nil {
		shared.OError(id, err.Error())
		wg.Done()
	}

	shared.OInfo(id, "Got pow challenge")

	w.Wait()

	shared.OInfo(id, "Got rec challenge")

	shared.OInfo(id, "Solving challenges ")

	err = manager.SolveRecaptchaChallange()
	if err != nil {
		shared.OError(id, err.Error())
		wg.Done()
		return
	}

	err = manager.SolvePowChallange()
	if err != nil {
		shared.OError(id, err.Error())
		wg.Done()
		return
	}

	queueId, err := manager.EnterQueue()
	if err != nil {
		shared.OError(id, err.Error())
		wg.Done()
		return
	}
	if queueId == "" {
		shared.OError(id, "Failed entering queue")
		wg.Done()
		return
	}

	shared.OInfo(id, "Entered queue, id: "+queueId)

	manager.QueueId = queueId

	finished := false
	old := 0
	interval := 2000
	for !finished {
		finished, old = manager.CheckStatus(queueId)
		if old > 10 {
			interval = old
		}

		time.Sleep(time.Millisecond * time.Duration(interval))
	}

	wg.Done()
	return
}

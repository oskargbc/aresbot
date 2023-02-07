package tools

import (
	models2 "aresbot/internal/aio/models"
	"aresbot/internal/aio/shared"
	l "aresbot/pkg/logger"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type ToolRunner struct {
	Tasks          []models2.Task
	Proxys         []models2.Proxy
	Settings       models2.Settings
	WebhookHandler shared.WebhookHandler
}

func (s *ToolRunner) Run(t string, extras map[string]interface{}) {
	rotator := shared.ProxyRotator{Proxys: s.Proxys}

	tool := AllTools[t]

	ToolRunner_(tool, s.Settings, rotator, s.WebhookHandler, t, extras)
}

func ToolRunner_(t Tool, settings models2.Settings, rotator shared.ProxyRotator, handler shared.WebhookHandler, name string, extras map[string]interface{}) {
	id := 1

	if !t.IsActive() {
		shared.OError(id, "Tool is not loaded..")
		time.Sleep(1 * time.Second)
		return
	}

	if t.NeedBackend() {
		resp, err := http.Get("https://www.firestormbot.com/api/check")
		if err != nil {
			shared.OError(id, "Can't connect with backend..")
			time.Sleep(1 * time.Second)
		}
		body, _ := ioutil.ReadAll(resp.Body)

		var response shared.ApiStatusResponse
		_ = json.Unmarshal(body, &response)

		if !response.Online {
			shared.OError(id, "Backend offline but is needed..")
			time.Sleep(1 * time.Second)
		}
	}

	err := t.GetSettings(settings, id)
	if err != nil {
		shared.OError(id, "Error while loading Settings (captcha provider)")
		l.ErrorLogger.Println("error while loading tool settings ", err)
		return
	}
	t.Run(rotator, extras, handler)
}

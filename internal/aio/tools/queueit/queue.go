package queueit

import (
	"aresbot/internal/aio/shared"
	"aresbot/pkg/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (s *QueueIt) GetEnterChallengeCookie(t string) {
	base := "https://%s.queue-it.net/?c=%s&e=%s&t=%s"

	url := fmt.Sprintf(base, s.Shop, s.C, s.E, t)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", s.Shop+".queue-it.net")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"98\", \"Google Chrome\";v=\"98\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("sec-fetch-site", "none")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("accept-language", "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7")

	res, err := s.Client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	logger.InfoLogger.Println("Get challenge: ", res.StatusCode)

	s.Client.Jar.SetCookies(res.Request.URL, res.Cookies())

	for _, v := range res.Cookies() {
		s.challengeCookie = strings.Split(v.Value, "u=")[1]
	}

	logger.InfoLogger.Println("Cookie: ", s.challengeCookie)
}

func (s *QueueIt) EnterQueue() (string, error) {
	base := "https://%s.queue-it.net/spa-api/queue/%s/%s/enqueue"
	url := fmt.Sprintf(base, s.C, s.C, s.E)
	method := "POST"

	r := s.solvedReC
	p := s.solvedPow

	ReChallenge := ChallengeSessions{
		SessionID:     r.SessionInfo.SessionID,
		Timestamp:     r.Timestamp,
		Checksum:      r.SessionInfo.Checksum,
		SourceIP:      r.SessionInfo.SourceIP,
		ChallengeType: r.SessionInfo.ChallengeType,
		Version:       6,
	}
	PowChallenge := ChallengeSessions{
		SessionID:     p.SessionInfo.SessionID,
		Timestamp:     p.Timestamp,
		Checksum:      p.SessionInfo.Checksum,
		SourceIP:      p.SessionInfo.SourceIP,
		ChallengeType: p.SessionInfo.ChallengeType,
		Version:       6,
	}

	challenges := []ChallengeSessions{}
	challenges = append(challenges, ReChallenge)
	challenges = append(challenges, PowChallenge)

	data := QueueItEnterRequest{
		ChallengeSessions: challenges,
		LayoutName:        s.LayoutName,
		CustomURLParams:   "",
		TargetURL:         s.TargetUrl,
		Referrer:          "https://www.google.com/",
		LayoutVersion:     s.LayoutName,
	}

	payload, _ := json.Marshal(data)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		logger.ErrorLogger.Println(err)
		return "", err
	}

	req.Header.Add("authority", s.C+".queue-it.net")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"98\", \"Google Chrome\";v=\"98\"")
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("origin", "https://"+s.C+".queue-it.net")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("accept-language", "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7")

	res, err := s.Client.Do(req)
	logger.InfoLogger.Println("Enterqueue: ", res.StatusCode)

	if err != nil {
		logger.ErrorLogger.Println(err)
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return "", err
	}

	logger.InfoLogger.Println(string(body))

	var resp QueueEnterResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return "", err
	}

	return resp.QueueID, nil
}

func (s *QueueIt) CheckStatus(QueueId string) (bool, int) {
	Timestamp := GenerateTimestamp()

	url := "https://" + s.C + ".queue-it.net/spa-api/queue/" + s.C + "/" + s.E + "/" + QueueId + "/status?sets=" + Timestamp
	method := "POST"
	payload := strings.NewReader(`{"targetUrl":"` + s.TargetUrl + `","customUrlParams":"","layoutVersion": "` + s.LayoutVersion + `","layoutName":"` + s.LayoutName + `","isClientRedayToRedirect":true,"isBeforeOrIdle":false}`)

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		logger.ErrorLogger.Println(err)
		return false, 0
	}
	req.Header.Add("authority", s.C+".queue-it.net")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"98\", \"Google Chrome\";v=\"98\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("origin", "https://"+s.C+".queue-it.net")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://"+s.C+".queue-it.net/?c="+s.C+"&e="+s.E)
	req.Header.Add("accept-language", "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7")

	res, err := s.Client.Do(req)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return false, 0
	}

	logger.InfoLogger.Println("Status Queue: ", res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return false, 0
	}
	var resp QueueItReponse
	_ = json.Unmarshal(body, &resp)
	logger.InfoLogger.Println(resp)

	if resp.ForecastStatus == "" {
		var respFinished QueueFinishedResponse
		_ = json.Unmarshal(body, &respFinished)

		if respFinished.IsRedirectToTarget {
			shared.OSuccess(s.TaskId, "Passed Queue, redirecting.. ")
			go shared.OpenBrowser(s.TaskId, respFinished.RedirectURL)
			s.SendWebhooks(respFinished.RedirectURL, s.QueueUrl)
			return true, 0
		}

		shared.OInfo(s.TaskId, "Status: "+resp.ForecastStatus+", Progress: "+strconv.Itoa(resp.Ticket.Progress))
		return false, resp.UpdateInterval
	} else {

		shared.OInfo(s.TaskId, "Status: "+resp.ForecastStatus+", Progress: "+strconv.Itoa(resp.Ticket.Progress))
		return false, resp.UpdateInterval
	}
}

func (s *QueueIt) SendWebhooks(uri string, url string) {
	n := time.Now()
	privat := shared.Webhook{
		Content:   "",
		Username:  "Firestorm Tools",
		AvatarURL: nil,
		Tts:       false,
		Embeds: []shared.Embeds{
			{
				Title:       "Passed Queue",
				Description: "you got redirected on " + s.Shop,
				URL:         uri,
				Timestamp:   time.Now(),
				Color:       6946682,
				Footer: shared.Footer{
					Text:    "Firestorm " + s.Handler.Version,
					IconURL: nil,
				},
				Thumbnail: shared.Thumbnail{
					URL: "https://upload.wikimedia.org/wikipedia/commons/5/5f/Queue-it-icon.png?20200908104108",
				},
				Author: shared.Author{},
				Image:  shared.Image{},
				Fields: []shared.Fields{
					{
						Name:   "Time in Queue",
						Value:  n.Sub(s.Start).String(),
						Inline: false,
					},
					{
						Name:   "Customer Id",
						Value:  s.C,
						Inline: false,
					},
					{
						Name:   "Redirected",
						Value:  "true",
						Inline: false,
					},
					{
						Name:   "Queue Id",
						Value:  s.QueueId,
						Inline: false,
					},
				},
			},
		},
	}

	public := shared.Webhook{
		Content:   "",
		Username:  "Firestorm Tools",
		AvatarURL: nil,
		Tts:       false,
		Embeds: []shared.Embeds{
			{
				Title:       "Passed Queue",
				Description: "User passed queue-it on " + s.Shop,
				URL:         url,
				Timestamp:   time.Now(),
				Color:       6946682,
				Footer: shared.Footer{
					Text:    "Firestorm " + s.Handler.Version,
					IconURL: nil,
				},
				Thumbnail: shared.Thumbnail{
					URL: "https://upload.wikimedia.org/wikipedia/commons/5/5f/Queue-it-icon.png?20200908104108",
				},
				Author: shared.Author{},
				Image:  shared.Image{},
				Fields: []shared.Fields{
					{
						Name:   "Time in Queue",
						Value:  n.Sub(s.Start).String(),
						Inline: false,
					},
					{
						Name:   "Redirected",
						Value:  "true",
						Inline: false,
					},
				},
			},
		},
	}

	log := shared.Webhook{
		Content:   "",
		Username:  "Firestorm Tools",
		AvatarURL: nil,
		Tts:       false,
		Embeds: []shared.Embeds{
			{
				Title:       "User Pass",
				Description: "User " + s.Handler.HyperData.Email + " passed the queue",
				URL:         uri,
				Timestamp:   time.Now(),
				Color:       0,
				Footer:      shared.Footer{},
				Image:       shared.Image{},
				Thumbnail:   shared.Thumbnail{},
				Author:      shared.Author{},
				Fields: []shared.Fields{
					{
						Name:   "Id",
						Value:  s.QueueId,
						Inline: false,
					},
					{
						Name:   "Time in queue",
						Value:  n.Sub(s.Start).String(),
						Inline: false,
					},
					{
						Name:   "License",
						Value:  s.Handler.HyperData.Key,
						Inline: false,
					},
				},
			},
		},
	}

	go s.Handler.SendCustomWebhook(public, privat, log)
}

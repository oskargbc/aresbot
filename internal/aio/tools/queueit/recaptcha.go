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
	"sync"
)

func (s *QueueIt) GetRecaptchaChallange(wg *sync.WaitGroup) error {
	wg.Add(1)

	shared.OInfo(s.TaskId, "Waiting for solved recaptcha")
	result, err := s.Solver.SolveRecaptcha("https://www.recaptcha.net/recaptcha/api2/userverify?k="+s.SiteKey, "", s.SiteKey)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	shared.OInfo(s.TaskId, "Solved recaptcha")

	s.solvedCaptchaCode = result
	wg.Done()

	return nil
}

func (s *QueueIt) SolveRecaptchaChallange() error {

	data := RecaptchaSolveData{
		ChallengeType: "recaptcha-invisible",
		SessionID:     s.solvedCaptchaCode,
		CustomerID:    s.C,
		EventID:       s.E,
		Version:       6,
	}

	payload, _ := json.Marshal(data)

	url := "https://" + s.C + ".queue-it.net/challengeapi/verify"

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Host", s.C+".queue-it.net")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("origin", "https://"+s.C+".queue-it.net")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("accept-language", "en,en-CA;q=0.9")

	res, err := s.Client.Do(req)

	logger.InfoLogger.Println("Solving ReC challenge status: ", res.StatusCode)

	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	var ReC RecaptchaSolveResponse
	err = json.Unmarshal(body, &ReC)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	s.solvedReC = ReC

	verify := strconv.FormatBool(s.solvedReC.IsVerified)

	shared.OInfo(s.TaskId, "ReC challenge status: "+verify)

	return nil
}

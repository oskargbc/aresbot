package queueit

import (
	"aresbot/internal/aio/shared"
	"aresbot/pkg/logger"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"

	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func (s *QueueIt) GetPowChallenge() error {

	urll := "https://" + s.Shop + ".queue-it.net/challengeapi/pow/challenge/" + s.challengeCookie

	req, err := http.NewRequest(http.MethodPost, urll, nil)

	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}
	req.Header.Add("Host", s.Shop+".queue-it.net")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	req.Header.Add("powtag-userid", s.challengeCookie)
	req.Header.Add("powtag-eventid", s.E)
	req.Header.Add("powtag-customerid", s.C)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	req.Header.Add("accept", "*/*")
	req.Header.Add("origin", "https://"+s.C+".queue-it.net")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("accept-language", "en,en-CA;q=0.9")

	res, err := s.Client.Do(req)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	var powReponse PowResponse
	_ = json.Unmarshal(body, &powReponse)
	s.powResponse = powReponse
	return nil
}

func (s *QueueIt) SolvePowChallange() error {
	r := s.powResponse

	postfix, hash := getPowHash(r.Parameters.Input, r.Parameters.ZeroCount)

	sessionIdPow := PowSessionId{
		UserID:    s.challengeCookie,
		Meta:      r.Meta,
		SessionID: r.SessionID,
		Solution: Solution{
			Postfix: postfix,
			Hash:    hash,
		},
		Tags: []string{
			"powTag-CustomerId:" + s.C,
			"powTag-EventId:" + s.E,
			"powTag-UserId:" + s.challengeCookie,
		},
		Stats: Stats{
			Duration:       3343,
			Tries:          1,
			UserAgent:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36",
			Screen:         "1920 x 1080",
			Browser:        "Chrome",
			BrowserVersion: "98.0.4758.102",
			IsMobile:       false,
			Os:             "Macintosh",
			OsVersion:      "10_15_7",
			CookiesEnabled: true,
		},
		Parameters: Parameters{
			Input:     r.Parameters.Input,
			ZeroCount: r.Parameters.ZeroCount,
		},
	}

	jsonSessionIdPow, _ := json.Marshal(sessionIdPow)

	baseSession := base64.StdEncoding.EncodeToString(jsonSessionIdPow)

	powData := PowSolveData{
		ChallengeType: "proofofwork",
		SessionID:     baseSession,
		CustomerID:    s.C,
		EventID:       s.E,
		Version:       6,
	}

	payload, _ := json.Marshal(powData)

	url := "https://" + s.C + ".queue-it.net/challengeapi/verify"

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))

	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}
	req.Header.Add("Host", s.Shop+".queue-it.net")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("origin", "https://"+s.Shop+".queue-it.net")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("accept-language", "en,en-CA;q=0.9")

	res, err := s.Client.Do(req)

	logger.InfoLogger.Println("Solving pow ", res.StatusCode)

	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 504 {
		res, err = s.Client.Do(req)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return err
		}
		logger.ErrorLogger.Println(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	logger.InfoLogger.Println(string(body))

	var powResponse PowSolveResponse
	err = json.Unmarshal(body, &powResponse)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	s.solvedPow = powResponse

	shared.OInfo(s.TaskId, "Pow challenge status: "+strconv.FormatBool(powResponse.IsVerified))
	return nil
}

func getPowHash(input string, zeroCount int) (int, string) {
	zeros := strings.Repeat("0", zeroCount)
	for postfix := 0; ; postfix++ {
		str := input + strconv.Itoa(postfix)
		hash := sha256.New()
		hash.Write([]byte(str))
		encodedHash := hex.EncodeToString(hash.Sum(nil))
		if strings.HasPrefix(encodedHash, zeros) {
			return postfix, encodedHash
		}
	}
}

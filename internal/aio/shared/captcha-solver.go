package shared

import (
	"aresbot/internal/aio/constants"
	api2captcha "aresbot/pkg/2captcha"
	"errors"
	"github.com/nuveo/anticaptcha"
	"time"
)

type captchaSolver struct {
	key         string
	provider    string
	twoCaptcha  *api2captcha.Client
	antiCaptcha *anticaptcha.Client
}

type CaptchaSolver interface {
	SolveHCaptcha(url string, proxy string, key string) (result string, err error)
	SolveRecaptcha(url string, proxy string, key string) (result string, err error)
}

func NewCaptchaSolver(Key string, Type string) CaptchaSolver {
	switch Type {
	case "2captcha":
		return &captchaSolver{
			key:        Key,
			provider:   Type,
			twoCaptcha: api2captcha.NewClient(Key),
		}
	case "anticaptcha":
		return &captchaSolver{
			key:         Key,
			provider:    Type,
			antiCaptcha: &anticaptcha.Client{APIKey: Key},
		}
	}
	return nil
}

func (s *captchaSolver) SolveHCaptcha(url string, proxy string, key string) (result string, err error) {
	if s.provider == "" {
		return "", constants.ErrNoCaptchaProvider
	}

	switch s.provider {
	case "2captcha":
		return s.solveHCaptchaWithTwoCaptcha(url, proxy, key)
	case "anticaptcha":
		return s.solveHCaptchaWithAntiCaptcha(url, proxy, key)
	}

	return "", nil
}

func (s *captchaSolver) SolveRecaptcha(url string, proxy string, key string) (result string, err error) {
	if s.provider == "" {
		return "", constants.ErrNoCaptchaProvider
	}

	switch s.provider {
	case "2captcha":
		return s.solveRecaptchaWithTwoCaptcha(url, proxy, key)
	case "anticaptcha":
		return s.solveRecaptchaWithAntiCaptcha(url, proxy, key)
	}

	return "", nil
}

func (s *captchaSolver) solveHCaptchaWithTwoCaptcha(url string, proxy string, key string) (string, error) {
	payload := api2captcha.HCaptcha{
		SiteKey: key,
		Url:     url,
	}

	req := payload.ToRequest()

	if proxy != "" {
		proxyType, proxyPart := prepareProxy(proxy)
		req.SetProxy(proxyType, proxyPart)
	}

	code, err := s.twoCaptcha.Solve(req)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (s *captchaSolver) solveHCaptchaWithAntiCaptcha(url string, proxy string, key string) (string, error) {
	return "", errors.New("not implemented")
}

func (s *captchaSolver) solveRecaptchaWithTwoCaptcha(url string, proxy string, key string) (string, error) {
	payload := api2captcha.ReCaptcha{
		SiteKey:   key,
		Url:       url,
		Invisible: true,
		Action:    "verify",
	}

	req := payload.ToRequest()

	if proxy != "" {
		proxyType, proxyPart := prepareProxy(proxy)
		req.SetProxy(proxyType, proxyPart)
	}

	code, err := s.twoCaptcha.Solve(req)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (s *captchaSolver) solveRecaptchaWithAntiCaptcha(url string, proxy string, key string) (string, error) {
	key, err := s.antiCaptcha.SendRecaptcha(url, key, 10*time.Second)
	if err != nil {
		return "", nil
	}

	return key, nil
}

func prepareProxy(proxy string) (string, string) {
	return "", ""
}

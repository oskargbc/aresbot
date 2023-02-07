package queueit

import (
	"aresbot/internal/aio/shared"
	"net/http"
	"strconv"
	"time"
)

func NewQueueItManager(timestamp, c, e, shop, sitekey, target, layoutName, layoutVersion string, client *http.Client, solver shared.CaptchaSolver, id int, handler shared.WebhookHandler, queueUrl string) QueueIt {
	return QueueIt{
		QueueId:           "",
		Timestamp:         timestamp,
		C:                 c,
		E:                 e,
		Shop:              shop,
		Solver:            solver,
		challengeCookie:   "",
		powResponse:       PowResponse{},
		solvedCaptchaCode: "",
		Client:            client,
		solvedPow:         PowSolveResponse{},
		solvedReC:         RecaptchaSolveResponse{},
		SiteKey:           sitekey,
		TargetUrl:         target,
		TaskId:            id,
		LayoutName:        layoutName,
		LayoutVersion:     layoutVersion,
		Handler:           handler,
		Start:             time.Now(),
		QueueUrl:          queueUrl,
	}
}

type QueueIt struct {
	QueueId   string
	Timestamp string
	C         string
	E         string
	Shop      string
	Solver    shared.CaptchaSolver
	Handler   shared.WebhookHandler

	challengeCookie string

	powResponse       PowResponse
	solvedCaptchaCode string

	Client *http.Client

	solvedPow PowSolveResponse
	solvedReC RecaptchaSolveResponse
	SiteKey   string
	TargetUrl string

	TaskId        int
	LayoutName    string
	LayoutVersion string
	Start         time.Time
	QueueUrl      string
}

func GenerateTimestamp() string {
	t := strconv.FormatInt(time.Now().Unix(), 10)

	return t
}

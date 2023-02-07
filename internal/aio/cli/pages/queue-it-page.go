package pages

import (
	"aresbot/internal/aio/cli/utils"
	"aresbot/internal/aio/constants"
	"aresbot/internal/aio/shared"
	"strconv"
	"strings"
)

func NewQueueItPage(d *shared.CliData) Page {
	return &QueueItPage{
		Data:  d,
		Route: "QueueItPage",
	}
}

type QueueItPage struct {
	Page

	Data  *shared.CliData
	Route string
}

func (s *QueueItPage) Load() (*string, shared.CliResponse) {
	utils.Clear()
	goTo := "Menu"
	utils.PrintHeader(s.Data.Version)

	queueUrl := utils.AwaitTextInput("QueueIt", "Enter queue-url: ", 9)
	if queueUrl == nil {
		return &s.Route, shared.CliResponse{}
	}

	entrys := utils.AwaitTextInput("QueueIt", "Enter queue entrys: ", 0)
	if entrys == nil {
		return &s.Route, shared.CliResponse{}
	}

	entrysInt, err := strconv.Atoi(*entrys)
	if err != nil {
		return &s.Route, shared.CliResponse{}
	}

	useProxy := false
	withProxy := utils.AwaitTextInput("QueueIt", "Use proxys y(es)/n(o):  ", 0)
	if withProxy == nil {
		return &s.Route, shared.CliResponse{}
	}
	if strings.ToLower(*withProxy) == "y" {
		useProxy = true
	} else {
		useProxy = false
	}

	data := map[string]interface{}{
		"useProxy": useProxy,
		"entrys":   entrysInt,
		"queueUrl": queueUrl,
	}

	a := constants.CmdStartTool
	resp := shared.CliResponse{
		Cmd:      &a,
		ToolName: "QueueIt",
		Extras:   data,
	}
	again := true

	for again {
		again, goTo, resp = utils.Start(s.Route, resp)
	}

	return &goTo, resp
}

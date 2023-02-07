package pages

import (
	"aresbot/internal/aio/cli/utils"
	"aresbot/internal/aio/constants"
	"aresbot/internal/aio/shared"
)

// new page

func NewMobilePage(d *shared.CliData) Page {
	return &MobilePage{
		Data:  d,
		Route: "MobilePage",
	}
}

type MobilePage struct {
	Page

	Data  *shared.CliData
	Route string
}

func (s *MobilePage) Load() (*string, shared.CliResponse) {
	cmd := constants.CmdStartTool
	goTo := "Menu"

	utils.PrintWithTag("MOBILE", "Connecting..")

	return &goTo, shared.CliResponse{
		ToolName: "Mobile",
		Cmd:      &cmd,
	}
}

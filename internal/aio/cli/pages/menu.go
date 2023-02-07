package pages

import (
	"aresbot/internal/aio/cli/utils"
	"aresbot/internal/aio/constants"
	"aresbot/internal/aio/shared"
)

func NewMenuPage(d *shared.CliData) Page {
	return &Menu{
		Data:  d,
		Route: "Menu",
	}
}

type Menu struct {
	Page

	Data  *shared.CliData
	Qt    bool
	Route string
}

func (s *Menu) Load() (*string, shared.CliResponse) {
	utils.Clear()
	utils.PrintHeader(s.Data.Version)
	s.menuHead()

	qt := "Start"
	if s.Qt {
		qt = "Stop"
	}

	choose := []string{"1", "2", "3", "4", "5", "6"}
	utils.PrintWithTag("MENU", "1 - Start task-group from .csv")
	utils.PrintWithTag("MENU", "2 - Start task-group by shop")
	utils.PrintWithTag("MENU", "3 - Reload data")
	utils.PrintWithTag("MENU", "4 - "+qt+" quicktask server")
	utils.PrintWithTag("MENU", "6 - Discord monitor")

	in := utils.AwaitInput("MENU", "Choose action: ", choose)

	var goTo string

	if in == nil {
		return &s.Route, shared.CliResponse{}
	}

	switch *in {
	case "1":
		goTo = "TaskGroupCsv"
	case "2":
		goTo = "TaskGroupShop"
	case "3":
		a := constants.CmdReload
		return &s.Route, shared.CliResponse{Cmd: &a}
	case "4":
		a := constants.CmdToggleQuicktasks
		return &s.Route, shared.CliResponse{Cmd: &a}
	case "5":
		a := constants.CmdStartDiscordMonitor
		return &s.Route, shared.CliResponse{Cmd: &a}
	}
	return &goTo, shared.CliResponse{}
}

func (s *Menu) menuHead() {
	s.Qt = constants.WithQuickTask
	bd := false

	/*
		r, err := http.Get("https://www.firestormbot.com/api/check")
		if err != nil {
			l.ErrorLogger.Println(err)
			bd = false
		}
		body, _ := ioutil.ReadAll(r.Body)
		var backendUp shared.ApiStatusResponse
		_ = json.Unmarshal(body, &backendUp)

		if !backendUp.Online {
			bd = false
		}
	*/
	// 1 section
	utils.PrintTag("MENU")
	utils.White.Print("Welcome ")
	utils.Yellow.Print(s.Data.Hyper.User.(map[string]interface{})["username"])
	utils.White.Print(" | Quicktask server: ")
	if s.Qt {
		utils.Green.Print("ONLINE")
	} else {
		utils.Red.Print("OFFLINE")
	}
	utils.White.Print(" | ")
	utils.White.Print("License: ")
	utils.Yellow.Println(s.Data.Hyper.Plan.Name)

	// 2 section
	utils.PrintTag("MENU")

	/*
		utils.White.Print("Database: ")
		if db {
			utils.Green.Print("CONNECTED")
		} else {
			utils.Red.Print("DISCONNECTED")
		}*/

	utils.White.Print("Backend: ")
	if bd {
		utils.Green.Print("ONLINE")
	} else {
		utils.Red.Print("OFFLINE")
	}

	utils.StaticClearLine()
}

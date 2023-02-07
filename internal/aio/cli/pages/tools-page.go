package pages

import (
	"aresbot/internal/aio/cli/utils"
	"aresbot/internal/aio/shared"
	"fmt"
	"strconv"
)

func NewToolsPage(d *shared.CliData) Page {
	return &ToolsPage{
		Data:  d,
		Route: "ToolsPage",
	}
}

type ToolsPage struct {
	Page

	Data  *shared.CliData
	Route string
}

func (s *ToolsPage) Load() (*string, shared.CliResponse) {
	utils.Clear()
	goTo := "Menu"
	utils.PrintHeader(s.Data.Version)

	inArr := []string{}

	allTools := strconv.Itoa(len(s.Data.Tools))

	utils.PrintWithTag("TOOLS", allTools+" Tools loaded")

	for k, v := range s.Data.Tools {
		kStr := strconv.FormatInt(int64(k+1), 10)

		inArr = append(inArr, kStr)

		utils.PrintTag("TOOLS")
		utils.White.Println(k+1, "- "+v)
	}
	utils.PrintWithTag("MENU", "[B] Go back")
	input := utils.AwaitInput("TASKS", "Choose tool: ", inArr)

	if input == nil {
		return &s.Route, shared.CliResponse{}
	}
	if *input == "B" {
		return &goTo, shared.CliResponse{}
	}

	switch *input {
	case "1":
		goTo = "QueueItPage"
		return &goTo, shared.CliResponse{}
	case "2":
		goTo = "MobilePage"
		return &goTo, shared.CliResponse{}
	}

	fmt.Println(*input)

	return &s.Route, shared.CliResponse{}
}

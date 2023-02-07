package pages

import (
	"aresbot/internal/aio/cli/utils"
	"aresbot/internal/aio/constants"
	"aresbot/internal/aio/shared"
	l "aresbot/pkg/logger"
	"strconv"
)

func NewTaskGroupCsvPage(d *shared.CliData) Page {
	return &TaskGroupCsv{
		Data:  d,
		Route: "TaskGroupCsv",
	}
}

type TaskGroupCsv struct {
	Page

	Data  *shared.CliData
	Route string
}

func (s *TaskGroupCsv) Load() (*string, shared.CliResponse) {
	utils.Clear()
	goTo := "Menu"
	utils.PrintHeader(s.Data.Version)

	var tasks []string
	var inArr []string

	for _, v := range s.Data.Groups {
		tasks = append(tasks, v.Name)
	}

	for k, v := range tasks {
		k++
		kStr := strconv.FormatInt(int64(k), 10)

		inArr = append(inArr, kStr)

		utils.PrintTag("TASKS")
		utils.White.Println(k, "- "+v)
	}

	utils.PrintWithTag("MENU", "[B] Go back")
	input := utils.AwaitInput("TASKS", "Choose taskgroup: ", inArr)
	utils.StaticClearLine()

	if input == nil {
		return &s.Route, shared.CliResponse{}
	}
	if *input == "B" {
		return &goTo, shared.CliResponse{}
	}

	i, _ := strconv.Atoi(*input)
	name := tasks[i-1]

	for _, v := range s.Data.Groups {
		if v.Name == name {
			again := true
			a := constants.CmdRunTasks
			resp := shared.CliResponse{
				Tasks: v.Tasks,
				Cmd:   &a,
			}
			for again {
				again, goTo, resp = utils.Start(s.Route, resp)
			}
			return &goTo, resp
		}
	}

	l.WarningLogger.Println("no tasks loaded")
	return &s.Route, shared.CliResponse{}
}

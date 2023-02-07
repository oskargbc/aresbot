package pages

import (
	"aresbot/internal/aio/cli/utils"
	"aresbot/internal/aio/constants"
	"aresbot/internal/aio/models"
	"aresbot/internal/aio/shared"
	l "aresbot/pkg/logger"
	"strconv"
)

func NewTaskGroupShopPage(d *shared.CliData) Page {
	return &TaskGroupShop{
		Data:  d,
		Route: "TaskGroupShop",
	}
}

type TaskGroupShop struct {
	Page

	Data  *shared.CliData
	Route string
}

func (s *TaskGroupShop) Load() (*string, shared.CliResponse) {
	utils.Clear()
	goTo := "Menu"
	utils.PrintHeader(s.Data.Version)

	loadedTaskPerShop := map[string]int{}
	inArr := []string{}
	ShopPerIndex := map[int]string{}

	for _, v := range s.Data.Shops {
		loadedTaskPerShop[v] = 0
	}

	for _, v := range s.Data.Groups {
		if len(v.Tasks) > 0 {
			for _, v := range v.Tasks {
				loadedTaskPerShop[v.Store]++
			}
		}
	}

	i := 0
	for k, v := range loadedTaskPerShop {
		i++
		ShopPerIndex[i] = k

		kStr := strconv.FormatInt(int64(i), 10)

		inArr = append(inArr, kStr)

		utils.PrintTag("TASKS")
		utils.White.Println(i, "- "+"[", v, "] "+k)
	}

	utils.PrintWithTag("MENU", "[B] Go back")
	input := utils.AwaitInput("TASKS", "Choose shop: ", inArr)

	if input == nil {
		return &s.Route, shared.CliResponse{}
	}
	if *input == "B" {
		return &goTo, shared.CliResponse{}
	}
	i, _ = strconv.Atoi(*input)
	name := ShopPerIndex[i]

	shopTasks := []models.Task{}

	for _, v := range s.Data.Groups {
		for _, v := range v.Tasks {
			if v.Store == name {
				shopTasks = append(shopTasks, v)
			}
		}
	}

	if len(shopTasks) < 1 {
		l.WarningLogger.Println("no shoptasks loaded")
	}

	again := true
	a := constants.CmdRunTasks
	resp := shared.CliResponse{
		Tasks: shopTasks,
		Cmd:   &a,
	}
	for again {
		again, goTo, resp = utils.Start(s.Route, resp)
	}
	return &goTo, resp
}

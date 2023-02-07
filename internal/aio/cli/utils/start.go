package utils

import "aresbot/internal/aio/shared"

func Start(last string, response shared.CliResponse) (bool, string, shared.CliResponse) {
	Clear()
	PrintWithTag("MENU", "Loaded tasks")
	PrintWithTag("MENU", "[S] Start")
	PrintWithTag("MENU", "[B] Go back")
	input := AwaitInput("MENU", "Choose action: ", []string{"S"})

	if input == nil {
		return true, "", shared.CliResponse{}
	} else if *input == "B" {
		return false, last, shared.CliResponse{}
	} else if *input == "S" {
		return false, "Menu", response
	}

	return true, "", shared.CliResponse{}
}

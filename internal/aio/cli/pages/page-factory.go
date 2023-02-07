package pages

import (
	"aresbot/internal/aio/shared"
)

type Page interface {
	Load() (*string, shared.CliResponse)
}

func LoadPages(d shared.CliData) map[string]Page {
	return map[string]Page{
		"Menu":          NewMenuPage(&d),
		"TaskGroupCsv":  NewTaskGroupCsvPage(&d),
		"TaskGroupShop": NewTaskGroupShopPage(&d),
		"ToolsPage":     NewToolsPage(&d),
		"QueueItPage":   NewQueueItPage(&d),
		"MobilePage":    NewMobilePage(&d),
	}
}

/*
// new page


func NewToolsPage(d *shared.CliData, db *sql.DB) Page {
	return &ToolsPage{
		Data:  d,
		Db:    db,
		Route: "ToolsPage",
	}
}

type ToolsPage struct {
	Page

	Data  *shared.CliData
	Db    *sql.DB
	Route string
}

func (s *ToolsPage) Load() (*string, shared.CliResponse) {
	return &s.Route, shared.CliResponse{}
}
*/

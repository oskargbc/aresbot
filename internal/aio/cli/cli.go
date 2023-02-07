package cli

import (
	"aresbot/internal/aio/cli/pages"
	"aresbot/internal/aio/shared"
	"database/sql"
	"fmt"
)

func NewCli(data shared.CliData) Cli {
	return Cli{
		Data: data,
	}
}

type Cli struct {
	Data  shared.CliData
	Db    *sql.DB
	Start *string
}

func (s *Cli) Handler(goTo *string) shared.CliResponse {
	/*
		executablePath, _ := osext.ExecutableFolder()
		os.Chdir(executablePath)
		db, err := sql.Open("sqlite3", "./FirestormDB.db")
		if err != nil {
			logger.ErrorLogger.Println(err)
			shared.OErrorF("Database connection lost")
			time.Sleep(time.Millisecond * 1500)
			os.Exit(-1)
		}*/
	s.Db = &sql.DB{}
	s.Start = goTo
	r := s.Router()

	return r
}

func (s *Cli) Routing() {

}

func (s *Cli) Router() shared.CliResponse {
	allPages := pages.LoadPages(s.Data)
	goTo := *s.Start
	fmt.Println(goTo)
	for {
		p := allPages[goTo]

		g, resp := p.Load()

		if resp.Cmd != nil {
			return resp
		}

		goTo = *g
	}
}

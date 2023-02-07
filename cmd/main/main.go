package main

import (
	"aresbot/internal/aio"
	bot_api "aresbot/internal/bot-api"
	"database/sql"
	"github.com/kardianos/osext"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	b := aio.NewBot()
	a := bot_api.NewApi()

	go a.Up()
	b.Run()
}

func test() {
	executablePath, _ := osext.ExecutableFolder()
	os.Chdir(executablePath)

	db, err := sql.Open("sqlite3", "./FirestormDB.db")
	if err != nil {
		log.Fatalln(err)
	}
	//var actions []*sql.Stmt
	//var createQuerys []string

	c, err := db.Prepare("CREATE TABLE settings (key_ TEXT PRIMARY KEY, value TEXT);")
	if err != nil {
		if err.Error() != "table settings already exists" {
			log.Fatalln(err)
		}
	} else {
		createDatabase(c, db)
	}

	db.Close()
}

func createDatabase(c *sql.Stmt, db *sql.DB) {
	_, err := c.Exec()
	if err != nil {
		log.Fatalln(err)
	}

	var actions []*sql.Stmt
	createQuerys := []string{
		"INSERT INTO settings (key_, value) VALUES (\"test\", \"true\");",
		"INSERT INTO settings (key_, value) VALUES (\"qt\", \"true\");",
	}
	for _, q := range createQuerys {
		stmt, err := db.Prepare(q)
		if err != nil {
			log.Fatalln(err)
		}
		actions = append(actions, stmt)
	}

	for _, action := range actions {
		_, err = action.Exec()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

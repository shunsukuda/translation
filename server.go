package main

import (
	"database/sql"
	"os"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

const (
	EIJIDB = `./eiji-141.db`
)

func ExecDB(dbfile string) {
	if _, err := os.Stat(dbfile); err != nil {
		return
	}
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		panic(err)
	}
	defer func() {
		stmt.Close()
		db.Close()
	}()
}

func getWord(c echo.Context) error {

}

func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":3300"))
}

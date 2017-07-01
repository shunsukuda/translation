package main

import (
	"bufio"
	"database/sql"
	"encoding/xml"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/cheggaaa/pb.v1"
)

type Pdic struct {
	Records []Record `xml:"record"`
}

type Record struct {
	Word  string `xml:"word"`
	Trans string `xml:"trans"`
}

func Load() []byte {
	f, err := os.Open(`./EIJI-141.xml`)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fstat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	lines := make([]byte, 0, fstat.Size())
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text())...)
	}
	if serr := scanner.Err(); serr != nil {
		panic(err)
	}
	return lines
}

func Parse() *Pdic {
	var pdic Pdic
	lines := Load()
	if err := xml.Unmarshal(lines, &pdic); err != nil {
		panic(err)
	}
	return &pdic
}

func CreateDB() {
	dbfile := "./eiji-141.db"
	if _, err := os.Stat(dbfile); err == nil {
		os.Remove(dbfile)
	}
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS eiji(id INTEGER PRIMARY KEY AUTOINCREMENT, word VARCHAR(255), trans VARCHAR(255))`)
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare(`INSERT INTO eiji(word, trans) VALUES (?, ?) `)
	if err != nil {
		panic(err)
	}
	records := Parse().Records
	pbar := pb.StartNew(len(records))
	for _, e := range records {
		if _, err = stmt.Exec(e.Word, e.Trans); err != nil {
			panic(err)
		}
		pbar.Increment()
	}
	_, err = db.Exec(`CREATE INDEX id_index ON eiji(id)`)
	defer func() {
		stmt.Close()
		db.Close()
	}()
}

func main() {
	CreateDB()
}

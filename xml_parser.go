package main

import (
	"bufio"
	"encoding/xml"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Pdic struct {
	Records []Record `xml:"record"`
}

type Record struct {
	Word  string `xml:"word"`
	Trans string `xml:"trans"`
}

func Load() []byte {
	f, err := os.Open(`./test.xml`)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fstat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	lines := make([]byte, 0, fstat.Size())
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text())...)
	}
	if serr := scanner.Err(); serr != nil {
		log.Fatal(err)
	}
	return lines
}

func Parse() *Pdic {
	var pdic Pdic
	lines := Load()
	if err := xml.Unmarshal(lines, &pdic); err != nil {
		log.Fatal(err)
	}
	return &pdic
}

func main() {
	Parse()
}

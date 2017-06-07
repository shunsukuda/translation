package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Data struct {
	word  string `xml:"word"`
	trans string `xml:"trans"`
}

type Record struct {
	Data []Data
}

func Load() []string {
	f, err := os.Open(`./EIJI-141.xml`)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	lines := make([]string, 0, 8664751)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if serr := scanner.Err(); serr != nil {
		log.Fatal(err)
	}

	return lines
}

func Parse() {
	var dict Directory
	err := xml.Unmarshal([]byte(Load()), &dict)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	l := Load()
	fmt.Println(l)
}

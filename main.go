package main

import (
	"countWords/countwords"
	"flag"
	"fmt"
	"os"
)

func main() {

	textFile := flag.String("textfile", "", "nome do arquivo contendo texto a ser analisado")
	flag.Parse()
	datum, err := os.ReadFile(*textFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(datum)
	statistics := countwords.CountWords(text)
	fmt.Println(statistics)
}

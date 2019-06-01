package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	textFile := flag.String("f", "", "The name / path of text file to be parsed")
	flag.Parse()

	file, err := ioutil.ReadFile(*textFile)
	if err != nil {
		Exit(fmt.Sprintf("Failed to open file '%s'", *textFile))
	}
	r := string(file)
	fmt.Println(r)
}

func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	wordFile := flag.String("w", "", "The name/path of text file to be parsed for count of words.")
	flag.Parse()

	if *wordFile == "" {
		printUsage()
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(*wordFile)
	if err != nil {
		log.Fatal(err)
	}
	r := string(file)

	wc := wordCount(r)
	fmt.Println(wc)
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	fmt.Println("\t -w\t The name/path of text file to be parsed for count of words.")
}

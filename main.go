package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	textFile := flag.String("f", "", "The name/path of text file to be parsed")
	flag.Parse()

	if len(*textFile) == 0 {
		printUsage()
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(*textFile)
	if err != nil {
		log.Fatal(err)
	}
	r := string(file)

	wc := wordCount(r)
	cc := len(r)
	fmt.Println("Number of words      : ", wc)
	fmt.Println("Number of characters : ", cc)
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	fmt.Println("\t -f\t The name/path of text file to be parsed")
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	wordFile := flag.String("w", "", "absolute path to text file for count of words.")
	charFile := flag.String("c", "", "absolute path to text file for count of characters.")
	flag.Parse()

	if *wordFile == "" && *charFile == "" {
		printUsage()
		os.Exit(1)
	}

	if *wordFile != "" && *charFile != "" {
		if *wordFile == *charFile {
			file, err := ioutil.ReadFile(*wordFile)
			if err != nil {
				log.Fatal(err)
			}
			r := string(file)

			wc := wordCount(r)
			cc := len(r)
			fmt.Println(wc)
			fmt.Println(cc)
		} else {
			filew, err := ioutil.ReadFile(*wordFile)
			if err != nil {
				log.Fatal(err)
			}
			filec, err := ioutil.ReadFile(*charFile)
			if err != nil {
				log.Fatal(err)
			}
			rw := string(filew)
			rc := string(filec)

			wc := wordCount(rw)
			cc := len(rc)
			fmt.Println(wc)
			fmt.Println(cc)
		}
	} else if *wordFile != "" {
		file, err := ioutil.ReadFile(*wordFile)
		if err != nil {
			log.Fatal(err)
		}
		r := string(file)

		wc := wordCount(r)
		fmt.Println(wc)
	} else {
		file, err := ioutil.ReadFile(*charFile)
		if err != nil {
			log.Fatal(err)
		}
		r := string(file)

		cc := len(r)
		fmt.Println(cc)
	}
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	fmt.Println("\t -w\t absolute path to text file for count of words")
	fmt.Println("\t -c\t absolute path to text file for count of characters")
}

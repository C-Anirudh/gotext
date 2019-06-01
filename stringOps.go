package main

import (
	"regexp"
)

func wordCount(text string) int {
	re := regexp.MustCompile(`[\S]+`)
	results := re.FindAllString(text, -1)
	return len(results)
}

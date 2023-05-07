package cmd

import (
	"fmt"

	"os"
	"regexp"
)

type Detail int

const (
	Days Detail = iota
	Months
	Years
)

var datePattern *regexp.Regexp = regexp.MustCompile("(200[0-9]|20[1-4][0-9]|2050)[-]?(0[1-9]|1[0-2])[-]?(0[1-9]|[12][0-9]|3[01])")

func organize(detail Detail) {
	fmt.Println("detail is", detail)

	files, err := os.ReadDir("./")

	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, file := range files {
		if file.IsDir() { continue }

		if datePattern.MatchString( file.Name() ) {
			fmt.Println(file.Name())
		}

	}




}
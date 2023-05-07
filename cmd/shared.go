package cmd

import (
	"fmt"

	"os"
)

type Detail int

const (
	Days Detail = iota
	Months
	Years
)

func organize(detail Detail) {
	fmt.Println("detail is", detail)

	files, err := os.ReadDir("./")

	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}




}
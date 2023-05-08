package cmd

import (
	"fmt"

	"os"
	"path"
	"regexp"
	"strconv"
)

type Detail int

const (
	Years  Detail = iota
	Months
	Days
)

var datePattern *regexp.Regexp = regexp.MustCompile("(200[0-9]|20[1-4][0-9]|2050)[-]?(0[1-9]|1[0-2])[-]?(0[1-9]|[12][0-9]|3[01])")
var monthsNames = [12] string {"Januar", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"}


func organize(detail Detail) {
	cwd, _ := os.Getwd() // get CWD
	files, err := os.ReadDir( cwd )

	if err != nil {
		fmt.Println("Error while reading directory:", err)
	}

	for _, file := range files {
		if file.IsDir() { continue }
		if !datePattern.MatchString( file.Name() ) { continue }

		// create strings
		var date []string = datePattern.FindStringSubmatch( file.Name() ) // returns string slice {"2025-03-17", "2025", "03", "17"}
		year, month, day := date[1], date[2], date[3]

		var yearPath string = path.Join(cwd, year)
		monthNumber, _ := strconv.Atoi(month)
		var monthDirName string = fmt.Sprintf("%s-%s", month, monthsNames[monthNumber - 1] ) // cretae names like "02-Februar"
		var monthPath string = path.Join(cwd, year, monthDirName)
		var dayPath string  = path.Join(monthPath, day)

		// create year dir
		_, err := os.Stat(yearPath)
		if os.IsNotExist( err ) {
			os.Mkdir( yearPath, 0755) // 0755 are UNIX permissions that dont have an effect on windws
		}

		// create month dir
		if detail == Months || detail == Days {
			_, err := os.Stat(monthPath)
			if os.IsNotExist( err ) {
				os.Mkdir( monthPath, 0755)
			}
		}

		// create day dir
		if detail == Days {
			_, err := os.Stat(dayPath)
			if os.IsNotExist( err ) {
				os.Mkdir( dayPath, 0755)
			}
		}

		// move file
		var oldFilePath string = path.Join( cwd, file.Name() )
		var newFilePath string

		switch detail {
		case Years:
			newFilePath = path.Join( yearPath,  file.Name() )
		case Months:
			newFilePath = path.Join( monthPath, file.Name() )
		case Days:
			newFilePath = path.Join( dayPath,   file.Name() )
		default:
			newFilePath = path.Join( yearPath,  file.Name() )
		}

		os.Rename(oldFilePath, newFilePath)

	}




}
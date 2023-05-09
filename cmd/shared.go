package cmd

import (
	"fmt"
	// "sort"
	"strings"

	"os"
	"path"
	"regexp"
	"strconv"

	"github.com/rwcarlsen/goexif/exif"
)

type Detail int
const (
	Years  Detail = iota
	Months
	Days
)

type Source int
const (
	FileName  Source = iota
	EXIF
	ModifiedTime
)

type Subject int
const (
	EXIFLoc  Subject = iota
	Name
)

var datePattern *regexp.Regexp = regexp.MustCompile("(200[0-9]|20[1-4][0-9]|2050)[-]?(0[1-9]|1[0-2])[-]?(0[1-9]|[12][0-9]|3[01])")
var monthsNames = [12] string {"Januar", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"}


func organizeByDate(detail Detail, source Source) {
	cwd, _ := os.Getwd() // get CWD
	files, err := os.ReadDir( cwd )

	if err != nil {
		fmt.Println("Error while reading directory:", err)
	}

	for _, file := range files {
		if file.IsDir() { continue }

		// make a string from the source where the date came from
		var dateSourceString string

		switch source {
			case FileName:
				dateSourceString = file.Name()
				break

			case EXIF:
				img, err := os.Open( file.Name() )
				if err != nil { fmt.Println("Error while opening image", err) }

				exifData, err := exif.Decode(img)
				if err != nil { continue } // error on extracting exif data

				exifTime, err := exifData.DateTime()
				if err != nil { continue } // error on extracting time data
				dateSourceString = exifTime.Local().String()

				img.Close()
				break

			case ModifiedTime:
				fileInfo, err := os.Stat( file.Name() )
				if err != nil { fmt.Println("Error while getting file Info", err) }
				dateSourceString = fileInfo.ModTime().String()
				break

			default: continue
		}

		// search for date in the provided string that was created from the date source
		if !datePattern.MatchString( dateSourceString ) { continue }

		// create strings
		var date []string = datePattern.FindStringSubmatch( dateSourceString ) // returns string slice {"2025-03-17", "2025", "03", "17"}
		year, month, day := date[1], date[2], date[3]

		var yearPath string = path.Join(cwd, year)
		monthNumber, _ := strconv.Atoi(month)
		var monthDirName string = fmt.Sprintf("%s-%s", month, monthsNames[monthNumber - 1] ) // cretae names like "02-Februar"
		var monthPath string = path.Join(cwd, year, monthDirName)
		var dayPath string  = path.Join(monthPath, day)

		// create year dir
		_, err := os.Stat(yearPath)
		if os.IsNotExist( err ) {
			os.Mkdir( yearPath, 0755 ) // 0755 are UNIX permissions that dont have an effect on windws
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

func organizeByTopic(topics []string) {
	cwd, _ := os.Getwd() // get CWD
	files, err := os.ReadDir( cwd )

	if err != nil {
		fmt.Println("Error while reading directory:", err)
	}


	for _, topic := range topics {
		for _, file := range files {
			if file.IsDir() { continue }

			var fileLower  string = strings.ToLower( file.Name() )
			var topicLower string = strings.ToLower( topic )

			// test if topic occurs in filename
			if !strings.Contains( fileLower, topicLower ) { continue }

			// create topic directory
			_, err := os.Stat( path.Join(cwd, topic) )
			if os.IsNotExist( err ) {
				os.Mkdir( path.Join(cwd, topic), 0755)
			}

			// move file
			var oldfilePath string = path.Join(cwd, file.Name())
			var newfilePath string = path.Join(cwd, topic, file.Name())

			os.Rename(oldfilePath, newfilePath)

		}
	}
}

func rename(source Source) {
	// TODO

}

func find(dirPath string, subject Subject, name string) {
	files, err := os.ReadDir( dirPath )

	if err != nil {fmt.Println("Error while reading directory:", err)}

	for _, file := range files {
		if file.IsDir() {
			// search subdirectories
			find( path.Join(dirPath, file.Name()), subject, name )

		} else {
			// look for subject

			switch subject {
				case EXIFLoc:
					img, err := os.Open( file.Name() )
					if err != nil { continue }

					exifData, err := exif.Decode(img)
					if err != nil { continue } // error on extracting exif data

					Lat, Long, err := exifData.LatLong()
					if err != nil { continue } // error on extracting time data
					img.Close()

					fmt.Println(file.Name(), Lat, Long)

				case Name:
					var fileLower string = strings.ToLower( file.Name() )
					var nameLower string = strings.ToLower( name )

					// test if topic occurs in filename
					if strings.Contains( fileLower, nameLower ) {
						fmt.Println(file.Name())
					}

				default: continue
			}
		}
	}
}
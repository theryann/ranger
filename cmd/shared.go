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

type Media int
const (
	All  Media = iota
	Extention
	Videos
	Music
	Pictures
)

var datePattern *regexp.Regexp = regexp.MustCompile("(200[0-9]|20[1-4][0-9]|2050)[-]?(0[1-9]|1[0-2])[-]?(0[1-9]|[12][0-9]|3[01])")
var monthsNames = [12] string {"Januar", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"}

var mediaFormats = map[string][]string {
	"image" : {"jpg", "jpeg","png", "tiff", "webp", "gif", "heic", "raw"},
	"video" : {"mp4", "mpeg", "avi", "mkv", "webm", "flv", "wmv", "asf", "m4v", "m2v", "3gp" },
	"music" : {"wav", "m4a", "aac", "pcm", "wma", "mp3", "ogg", "flac"},
}

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
				if err != nil { continue }

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
			if subject == Name {
				var fileLower string = strings.ToLower( file.Name() )
				var nameLower string = strings.ToLower( name )

				// test if topic occurs in directory name
				if strings.Contains( fileLower, nameLower ) {
					fmt.Println(file.Name())
				}
			}

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

func compile(dirPath string, rootPath string) {
	files, err := os.ReadDir( dirPath )

	if err != nil {fmt.Println("Error while reading directory:", err)}

	for _, file := range files {
		if file.IsDir() {
			// search subdirectories
			compile( path.Join(dirPath, file.Name()), rootPath )
			err := os.Remove( path.Join(dirPath, file.Name()) )
			if err != nil { fmt.Printf("err: %v\n", err) }

		} else {
			// move files to root
			var oldFilePath string = path.Join( dirPath,  file.Name() )
			var newFilePath string = path.Join( rootPath, file.Name() )
			os.Rename(oldFilePath, newFilePath)
		}
	}
}

func typify(media Media) {
	cwd, _ := os.Getwd() // get CWD
	files, err := os.ReadDir( cwd )
	if err != nil { fmt.Println("Error while reading directory:", err) }

	for _, file := range files {
		if file.IsDir() { continue }

		var extention string = strings.ToLower( path.Ext( file.Name() ) )  // strip point from extention
		if len(extention) > 0 {
			extention = extention[1:]
			if extention == "ini" { continue }  // skip desktop.ini file
		} else {
			continue
		}

		var formatIdentified bool = false
		var newFilePath string


		switch media {
			case All:
				// check all media formats
				for mediaType, formats := range mediaFormats {
					if formatIdentified { break }
					for _, format := range formats {
						if format != extention {
							continue
						}
						// create media directory
						_, err := os.Stat( path.Join(cwd, mediaType) )
						if os.IsNotExist( err ) {
							os.Mkdir( path.Join(cwd, mediaType), 0755)
						}
						newFilePath = path.Join(cwd, mediaType, file.Name())
						formatIdentified = true
						break
					}
				}
			case Extention:
				// create extention directory
				_, err := os.Stat( path.Join(cwd, extention) )
				if os.IsNotExist( err ) {
					os.Mkdir( path.Join(cwd, extention), 0755)
				}
				newFilePath = path.Join(cwd, extention, file.Name())
				formatIdentified = true
				break
			case Pictures:
				for _, format := range mediaFormats["image"] {
					if format != extention {
						continue
					}
					// create media directory
					_, err := os.Stat( path.Join(cwd, "images") )
					if os.IsNotExist( err ) {
						os.Mkdir( path.Join(cwd, "images"), 0755)
					}
					newFilePath = path.Join(cwd, "images", file.Name())
					formatIdentified = true
					break
				}
			case Music:
				for _, format := range mediaFormats["music"] {
					if format != extention {
						continue
					}
					// create media directory
					_, err := os.Stat( path.Join(cwd, "audio") )
					if os.IsNotExist( err ) {
						os.Mkdir( path.Join(cwd, "audio"), 0755)
					}
					newFilePath = path.Join(cwd, "audio", file.Name())
					formatIdentified = true
					break
				}
			case Videos:
				for _, format := range mediaFormats["video"] {
					if format != extention {
						continue
					}
					// create media directory
					_, err := os.Stat( path.Join(cwd, "video") )
					if os.IsNotExist( err ) {
						os.Mkdir( path.Join(cwd, "video"), 0755)
					}
					newFilePath = path.Join(cwd, "video", file.Name())
					formatIdentified = true
					break
				}

		}

		// in case any container has been found move file
		if formatIdentified {
			var oldFilePath string = path.Join(cwd, file.Name())
			os.Rename(oldFilePath, newFilePath)
		}

	}
}
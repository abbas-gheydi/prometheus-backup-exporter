package checkfiles

import (
	"fmt"
	"log"
	"os"
	"strings"
)
//Run checks files and get date and size and return its values
func Run(path string) (pathName string, size string, date string) {

	pathName = strings.ReplaceAll(path, `/`, "_")
	pathName = strings.ReplaceAll(pathName, `.`, "_")
	pathName = strings.ReplaceAll(pathName, `-`, "_")

	file, err := os.Stat(path)

	if err != nil {
		log.Print(err)

		size = "0"
		date = size
	} else {

		size = fmt.Sprintf("%v", float64(file.Size())/1024)
		ModTime := file.ModTime()
		date = fmt.Sprintf("%v", ModTime.Format("20060102150405"))
	}
	return

}


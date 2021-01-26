package utils

import (
	"log"
	"path"
	"path/filepath"
	"strings"
)

func getFile(pwd string) []string {
	filepathNames, err := filepath.Glob(filepath.Join(pwd, "*.mp3"))
	if err != nil {
		log.Fatal(err)
	}
	return filepathNames
}

func resFileName(i int) string {
	files := getFile("/home/pi/music/")
	if i == 0 {
		return files[0]
	}
	return files[i-1]
}

func formatFileName(file string) string {
	filenameWithSuffix := path.Base(file)
	fileSuffix := path.Ext(filenameWithSuffix)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}

package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"../entity"

	"github.com/subosito/gotenv"
)

var includedExtensions = []string{}
var excludedDirs = []string{}
var excludedFiles = []string{}

func init() {

	gotenv.Load()

	includedExtensions = strings.Split(os.Getenv("VALID_EXT"), ",")
	excludedDirs = strings.Split(os.Getenv("EXCLUDE_DIR"), ",")
	excludedFiles = strings.Split(os.Getenv("EXCLUDE_FILE"), ",")

}

// Hash gives hash value of a string
func Hash(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(h.Sum(nil))
}

// Contains check strings in an array
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// SuitableCheck controls path and extensions of file via excludeXXX and includeXXX variables
func SuitableCheck(path string, name string) bool {
	ext := filepath.Ext(path)
	if !Contains(includedExtensions, ext) {
		return false
	}

	if Contains(excludedDirs, name) {
		return false
	}

	if Contains(excludedFiles, name) {
		return false
	}

	return true
}

// Counts scans path and gives directory and file count for the path recursively
func Counts(scanPath string) (int, int) {
	var dirCount = 0
	var fileCount = 0
	err := filepath.Walk(scanPath,
		func(path string, info os.FileInfo, err error) error {

			if !SuitableCheck(path, info.Name()) {
				return nil
			}

			if info.IsDir() {
				dirCount++
			} else {
				fileCount++
			}

			return err
		})

	if err != nil {

		log.Println(err)
	}

	return dirCount, fileCount
}

//ComparesAsJSON dump map as JSON
func ComparesAsJSON(compares map[string][]entity.FilePath) string {
	jsonString, _ := json.MarshalIndent(compares, "", "    ")

	return string(jsonString)
	/*
		for key, paths := range compares {
			fmt.Printf("\n\n>>> %v\n", key)
			for _, path := range paths {
				fmt.Printf("\t%v\n", path.Path)
			}
		}*/
}

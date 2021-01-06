package scan

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Comparator struct {
	includedExtensions []string
	excludedDirs       []string
	excludedFiles      []string
}

func NewComparator(includedExtensions []string, excludedDirs []string, excludedFiles []string) *Comparator {
	return &Comparator{includedExtensions: includedExtensions, excludedDirs: excludedDirs, excludedFiles: excludedFiles}
}

// Contains check strings in an array
func (c *Comparator) contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// SuitableCheck controls path and extensions of file via excludeXXX and includeXXX variables
func (c *Comparator) suitableCheck(path string, name string) bool {
	ext := filepath.Ext(path)
	if !c.contains(c.includedExtensions, ext) {
		return false
	}

	if c.contains(c.excludedDirs, name) {
		return false
	}

	if c.contains(c.excludedFiles, name) {
		return false
	}

	return true
}

// Counts scans path and gives directory and file count for the path recursively
func (c *Comparator) Counts(scanPath string) (int, int) {
	var dirCount = 0
	var fileCount = 0
	err := filepath.Walk(scanPath,
		func(path string, info os.FileInfo, err error) error {

			if !c.suitableCheck(path, info.Name()) {
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
func (c *Comparator) ComparesAsJSON(compares map[string][]FilePath) string {
	jsonString, _ := json.MarshalIndent(compares, "", "    ")

	return string(jsonString)
}

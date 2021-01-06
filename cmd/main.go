package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/cheggaaa/pb"

	"github.com/kemalbayindir/dupfifi/scan"

	"github.com/subosito/gotenv"
)

// scanPath variable for scan
var scanPath = ""

var includedExtensions []string
var excludedDirs []string
var excludedFiles []string

var compares map[string][]scan.FilePath

func init() {

	gotenv.Load()

	includedExtensions = strings.Split(os.Getenv("VALID_EXT"), ",")
	excludedDirs = strings.Split(os.Getenv("EXCLUDE_DIR"), ",")
	excludedFiles = strings.Split(os.Getenv("EXCLUDE_FILE"), ",")
	scanPath = os.Getenv("SCAN_PATH")

	compares = make(map[string][]scan.FilePath)
}

func main() {
	comparator := scan.NewComparator(includedExtensions, excludedDirs, excludedFiles)

	fmt.Printf("\nScanning Folder \t: %v \n", scanPath)
	dirCount, fileCount := comparator.Counts(scanPath)
	fmt.Printf("Directory Count \t: %v \n", dirCount)
	fmt.Printf("File Count \t\t: %v \n\n", fileCount)

	bar := pb.StartNew(fileCount)
	compares, totalSize, err := comparator.Walk(scanPath, func() {
		bar.Increment()
	})

	if err != nil {
		log.Println(err)
	}

	fSize := float64(float64(totalSize/1024) / 1024)
	fmt.Printf("\n\n TOTAL SIZE : %.2f MB\n\n", fSize)

	fmt.Printf("\n LIST OF COMPARES")
	fmt.Printf("\n ********************** \n")
	jsonStr := comparator.ComparesAsJSON(compares)
	fmt.Println(jsonStr)

	//http.HandleFunc("/listdir", listdir)
	//http.ListenAndServe("localhost:8082", nil)
}

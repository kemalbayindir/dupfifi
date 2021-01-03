package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"strings"
	"time"

	"./entity"
	"./utils"

	"github.com/cheggaaa/pb"
	"github.com/subosito/gotenv"
)

// scanPath variable for scan
var scanPath = ""

var includedExtensions = []string{}
var excludedDirs = []string{}
var excludedFiles = []string{}

var compares map[string][]entity.FilePath

func init() {

	gotenv.Load()

	includedExtensions = strings.Split(os.Getenv("VALID_EXT"), ",")
	excludedDirs = strings.Split(os.Getenv("EXCLUDE_DIR"), ",")
	excludedFiles = strings.Split(os.Getenv("EXCLUDE_FILE"), ",")
	scanPath = os.Getenv("SCAN_PATH")

	compares = make(map[string][]entity.FilePath)
}

func main() {

	fmt.Printf("\nScanning Folder \t: %v \n", scanPath)
	dirCount, fileCount := utils.Counts(scanPath)
	fmt.Printf("Directory Count \t: %v \n", dirCount)
	fmt.Printf("File Count \t\t: %v \n\n", fileCount)

	var totalSize int64 = 0

	bar := pb.StartNew(fileCount)

	err := filepath.Walk(scanPath,
		func(path string, info os.FileInfo, err error) error {

			if !info.IsDir() && (info.Mode()&os.ModeSymlink == 0) {

				if !utils.SuitableCheck(path, info.Name()) {
					return nil
				}

				if err != nil {
					return err
				}

				hashValue := ""
				if !info.IsDir() {
					hashValue = utils.Hash(path)
				}

				totalSize = totalSize + info.Size()

				compares[hashValue] = append(compares[hashValue], entity.FilePath{Path: path})

				bar.Increment()
				time.Sleep(time.Millisecond)

			}
			// else {
			//	fmt.Println(info.Name())
			//}
			return err
		})

	if err != nil {
		log.Println(err)
	}

	fSize := float64(float64(totalSize/1024) / 1024)
	fmt.Printf("\n\n TOTAL SIZE : %.2f MB\n\n", fSize)

	fmt.Printf("\n LIST OF COMPARES")
	fmt.Printf("\n ********************** \n")
	jsonStr := utils.ComparesAsJSON(compares)
	fmt.Println(jsonStr)

	//http.HandleFunc("/listdir", listdir)
	//http.ListenAndServe("localhost:8082", nil)
}

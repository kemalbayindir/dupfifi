package main

import (
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof"
	"strings"

	"github.com/cheggaaa/pb"

	"github.com/kemalbayindir/dupfifi/scan"
)

func main() {

	flagScanPath := flag.String("scanPath", "./", "Target path to scan process")
	flagIncludedExtensions := flag.String("includedExtensions", ".png,.jpg,.jpeg,.bmp", "Allowed dile extension(s) during scan process. Please use comma to seperate extensions.")
	flagExcludedDirs := flag.String("excludedDirs", ".git,node_modules", "Not allowed directories to scan process. Please use comma to seperate extensions.")
	flagExcludedFiles := flag.String("excludedFiles", ".DS_Store", "Not allowed files to scan process. Please use comma to seperate extensions.")
	flag.Parse()

	includedExtensions := strings.Split(*flagIncludedExtensions, ",")
	excludedDirs := strings.Split(*flagExcludedDirs, ",")
	excludedFiles := strings.Split(*flagExcludedFiles, ",")
	scanPath := *flagScanPath

	//fmt.Println("Process Parameters")
	//fmt.Println("------------------")
	//fmt.Println("scanPath:", scanPath)
	//fmt.Println("includedExtensions:", includedExtensions)
	//fmt.Println("excludedDirs:", excludedDirs)
	//fmt.Println("excludedFiles:", excludedFiles)

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

package scan

import (
	"os"
	"path/filepath"
)

type TickerFunc func()

type Walker interface {
	Walk(scanPath string) (map[string][]FilePath, int64, error)
}

func (c *Comparator) Walk(scanPath string, tickerFunc TickerFunc) (map[string][]FilePath, int64, error) {

	compares := make(map[string][]FilePath)
	var totalSize int64 = 0
	err := filepath.Walk(scanPath,
		func(path string, info os.FileInfo, err error) error {

			if !info.IsDir() && (info.Mode()&os.ModeSymlink == 0) {

				if !c.suitableCheck(path, info.Name()) {
					return nil
				}

				if err != nil {
					return err
				}

				hashValue := ""
				if !info.IsDir() {
					hashValue = Hash(path)
				}

				totalSize = totalSize + info.Size()

				compares[hashValue] = append(compares[hashValue], FilePath{Path: path})
				tickerFunc()

				//time.Sleep(time.Millisecond)

			}

			return err
		})

	return compares, totalSize, err
}

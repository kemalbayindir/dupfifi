package scan

import (
	"os"
	"path/filepath"
)

// TickerFunc is an alias
type TickerFunc func()

// Walker interface represents mock ability to Walk operation
type Walker interface {
	Walk(scanPath string) (map[string][]FilePath, int64, error)
}

// Walk searches path and returns found files according to filters
func (c *Comparator) Walk(scanPath string, tickerFunc TickerFunc) (map[string][]FilePath, int64, error) {

	compares := make(map[string][]FilePath)
	var totalSize int64 = 0
	err := filepath.Walk(scanPath,
		func(fpath string, info os.FileInfo, err error) error {

			if !info.IsDir() && (info.Mode()&os.ModeSymlink == 0) {

				if !c.suitableCheck(fpath, info.Name()) {
					return nil
				}

				if err != nil {
					return err
				}

				hashValue := ""
				if !info.IsDir() {
					hashValue = Hash(fpath)
				}

				totalSize = totalSize + info.Size()

				compares[hashValue] = append(compares[hashValue], FilePath{Path: fpath})
				tickerFunc()

				//time.Sleep(time.Millisecond)

			}

			return err
		})

	return compares, totalSize, err
}

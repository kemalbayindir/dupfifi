package scan

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
)

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

package tools

import (
	"github.com/rs/xid"
	"io"
	"net/http"
	"os"
)

func GenerateId(codeObject string) string {
	newId := xid.New()
	return codeObject + newId.String()
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func OpenFile(fileName string) *os.File {
	xmlFile, err := os.Open(fileName)
	Check(err)
	return xmlFile
}

func DownloadFile(filePath string, url string) {
	// Create the file
	out, err := os.Create(filePath)
	Check(err)
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	Check(err)
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	Check(err)
}

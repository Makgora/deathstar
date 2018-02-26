package tools

import (
	"github.com/rs/xid"
	"os"
	"fmt"
	"net/http"
	"io"
)

func GenerateId(codeObject string) string {
	newId := xid.New()
	return codeObject + newId.String()
}

func OpenFile(fileName string) *os.File {
	xmlFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return xmlFile
}

func DownloadFile(filePath string, url string) error {

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

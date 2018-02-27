package tools_tests

import (
	"DeathStar/models/tools"
	"bytes"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestGenerateId(t *testing.T) {
	newId := tools.GenerateId("")
	if len(newId) != 20 {
		t.Error("Wrong size ! Expected 20 got ,", len(newId))
	}
	newId = tools.GenerateId("T")
	if len(newId) != 21 {
		t.Error("Wrong size ! Expected 21 got ,", len(newId))
	}
	if newId[0] != 'T' {
		t.Error("Wrong first char ! Expected T got ", string(newId[0]))
	}
}

func TestDownloadFile(t *testing.T) {
	url := "https://images.monaco-parkings.mc/places.xml"
	filePath := "placesT.xml"

	tools.DownloadFile(filePath, url)
	cmd := exec.Command("wget", url, "-O", "places.xml")
	err := cmd.Run()
	tools.Check(err)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Error("File doesn't exist !")
	}
	cmd = exec.Command("diff", "places.xml", filePath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	tools.Check(err)
	if out.String() != "" {
		t.Error("Wrong content ! Expected '' got '", out.String(), "'")
	}
	cmd = exec.Command("rm", "places.xml", filePath)
	err = cmd.Run()
	tools.Check(err)
}

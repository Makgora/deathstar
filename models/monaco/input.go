package monaco

import (
	"DeathStar/models/tools"
	"io/ioutil"
	"encoding/xml"
)

func UpdateMonaco() {
	filePath := "places.xml"
	url := "https://images.monaco-parkings.mc/places.xml"

	// Download places.xml
	tools.DownloadFile(filePath, url)
	// Open and read places.xml
	xmlData, err := ioutil.ReadFile(filePath)
	tools.Check(err)
	// Parse xmlData
	err = xml.Unmarshal(xmlData, )
	tools.Check(err)

}

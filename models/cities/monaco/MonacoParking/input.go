package MonacoParking

import (
	"DeathStar/models/tools"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

type presents struct {
	PresentsUpdate   string `xml:"presentsUpdate"`
	PresentsTotal    int    `xml:"presentsTotal"`
	PresentsAbonnes  int    `xml:"presentsAbonnes"`
	PresentsHoraires int    `xml:"presentsHoraires"`
}

type parc struct {
	NumParc            string   `xml:"numParc"`
	LibelleParc        string   `xml:"libelleParc"`
	StatusParc         string   `xml:"statusParc"`
	PlacesLibresParc   int      `xml:"placesLibresParc"`
	PlacesLibresUpdate string   `xml:"placesLibresUpdate"`
	Presents           presents `xml:"presents"`
}

type quartier struct {
	NomQuartier                   string `xml:"nomQuartier"`
	TotalLibresQuartier           string `xml:"totalLibresQuartier"`
	TotalPresentsQuartier         string `xml:"totalPresentsQuartier"`
	TotalPresentsHorairesQuartier string `xml:"totalPresentsHorairesQuartier"`
	TotalPresentsAbonnesQuartier  string `xml:"totalPresentsAbonnesQuartier"`
	Parc                          []parc
}

type xmlStruct struct {
	TotalPresent int    `xml:"totalPresent"`
	TotalLibre   int    `xml:"totalLibre"`
	LastUpdate   string `xml:"lastUpdate"`
	Quartier     []quartier
}

func dlParseXml() xmlStruct {
	filePath := "./models/cities/monaco/MonacoParking/places.xml"
	url := "https://images.monaco-parkings.mc/places.xml"

	// Download places.xml
	tools.DownloadFile(filePath, url)

	// Open and read places.xml
	xmlData, err := ioutil.ReadFile(filePath)
	tools.Check(err)

	// Delete first line to avoid encoding error
	xmlData = xmlData[44:]

	// Replace 'Obsolete' by O to allow int conversion
	xmlData = []byte(strings.Replace(string(xmlData), "Obsolete", "0", -1))

	// Parse xmlData
	s := xmlStruct{}
	err = xml.Unmarshal(xmlData, &s)
	tools.Check(err)
	return s
}

func PrintXmlStruct(s xmlStruct) {
	fmt.Printf("TotalPresent: %#v\n", s.TotalPresent)
	fmt.Printf("TotalLibre: %#v\n", s.TotalLibre)
	fmt.Printf("LastUpdate: %#v\n", s.LastUpdate)

	fmt.Printf("\nListe des Quartiers\n")
	for _, c := range s.Quartier {
		fmt.Printf("\n\tNomQuartier: %#v\n", c.NomQuartier)
		fmt.Printf("\tTotalLibresQuartier: %#v\n", c.TotalLibresQuartier)
		fmt.Printf("\tTotalPresentsQuartier: %#v\n", c.TotalPresentsQuartier)
		fmt.Printf("\tTotalPresentsHoraires: %#v\n", c.TotalPresentsHorairesQuartier)
		fmt.Printf("\tTotalPresentsAbonnes: %#v\n", c.TotalPresentsAbonnesQuartier)
		for _, z := range c.Parc {
			fmt.Printf("\n\t\tNumParc: %#v\n", z.NumParc)
			fmt.Printf("\t\tLibelleParc: %#v\n", z.LibelleParc)
			fmt.Printf("\t\tStatusParc: %#v\n", z.StatusParc)
			fmt.Printf("\t\tPlacesLibresParc: %#v\n", z.PlacesLibresParc)
			fmt.Printf("\t\tPlacesLibresUpdate: %#v\n", z.PlacesLibresUpdate)
			fmt.Printf("\t\tPresentsUpdate: %#v\n", z.Presents.PresentsUpdate)
			fmt.Printf("\t\tPresentsTotal: %#v\n", z.Presents.PresentsTotal)
			fmt.Printf("\t\tPresentsAbonnes: %#v\n", z.Presents.PresentsAbonnes)
			fmt.Printf("\t\tPresentsHoraires: %#v\n", z.Presents.PresentsHoraires)
		}
	}
}

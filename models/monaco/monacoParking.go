package monaco

import (
	"DeathStar/models/tools"
	"io/ioutil"
	"fmt"
	"encoding/xml"
	"DeathStar/models/objects"
)

type presents struct {
	PresentsUpdate   string 	`xml:"presentsUpdate"`
	PresentsTotal    int 		`xml:"presentsTotal"`
	PresentsAbonnes  string 	`xml:"presentsAbonnes"`
	PresentsHoraires string		`xml:"presentsHoraires"`
}

type parc struct {
	NumParc            int    `xml:"numParc"`
	LibelleParc        string `xml:"libelleParc"`
	StatusParc         string `xml:"statusParc"`
	PlacesLibresParc   int    `xml:"LibresParc"`
	PlacesLibresUpdate string 	`xml:"placesLibresUpdate"`
	Presents           presents `xml:"presents"`
}

type quartier struct {
	NomQuartier                   string `xml:"nomQuartier"`
	TotalLibresQuartier           string `xml:"totalLibresQuartier"`
	TotalPresentsQuartier         string `xml:"totalPresentsQuartier"`
	TotalPresentsHorairesQuartier string `xml:"totalPresentsHorairesQuartier"`
	TotalPresentsAbonnesQuartier  string `xml:"totalPresentsAbonnesQuartier"`
	Parc	                      []parc
}

type document struct {
	TotalPresent int    `xml:"totalPresent"`
	TotalLibre   int    `xml:"totalLibre"`
	LastUpdate   string `xml:"lastUpdate"`
	Quartier    []quartier
}

func updateMonacoParkings() {
	filePath := "places.xml"
	url := "https://images.monaco-parkings.mc/places.xml"

	// Download places.xml
	tools.DownloadFile(filePath, url)

	// Open and read places.xml
	xmlData, err := ioutil.ReadFile(filePath)
	tools.Check(err)

	// Delete first line to avoid encoding error
	xmlData = xmlData[44:]

	// Parse xmlData
	v := document{}
	err = xml.Unmarshal(xmlData, &v)
	tools.Check(err)

	// Populate liveDB

	objects.GetLiveDB()
}

func printXmlStruct(v document) {
	fmt.Printf("TotalPresent: %#v\n", v.TotalPresent)
	fmt.Printf("TotalLibre: %#v\n", v.TotalLibre)
	fmt.Printf("LastUpdate: %#v\n", v.LastUpdate)

	fmt.Printf("\nListe des Quartiers\n")
	for _, c := range v.Quartier {
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
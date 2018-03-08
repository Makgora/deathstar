package tools

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type district struct {
	Name string `xml:"Name"`
}

type DistrictsXml struct {
	XMLName   xml.Name `xml:"Districts"`
	Districts []district `xml:"District"`
}

func ParseDistrictsXml(filePath string) DistrictsXml {
	xmlData, err := ioutil.ReadFile(filePath)
	Check(err)

	s := DistrictsXml{}
	err = xml.Unmarshal(xmlData, &s)
	Check(err)
	return s
}

func PrintDistrictXml(districtsXml DistrictsXml) {
	fmt.Println("DistrictXml:")
	for _, v := range districtsXml.Districts {
		fmt.Println("\nName: " + v.Name)
	}
}

type parking struct {
	Id          string `xml:"Id"`
	Name        string `xml:"Name"`
	Owner       string `xml:"Owner"`
	District    string `xml:"District"`
	SpacesCount int    `xml:"SpacesCount"`
	Address     string `xml:"Address"`
	Description string `xml:"Description"`
}

type ParkingsXml struct {
	XMLName   xml.Name `xml:"Parkings"`
	Parkings []parking `xml:"Parking"`
}

func ParseParkingsXml(filePath string) ParkingsXml {
	xmlData, err := ioutil.ReadFile(filePath)
	Check(err)

	s := ParkingsXml{}
	err = xml.Unmarshal(xmlData, &s)
	Check(err)
	return s
}

func PrintParkingsXml(parkingsXml ParkingsXml) {
	for _, v := range parkingsXml.Parkings {
		fmt.Println("Id: " + v.Id)
		fmt.Println("Name: " + v.Name)
		fmt.Println("Owner: " + v.Owner)
		fmt.Println("District: " + v.District)
		fmt.Println("SpacesCount: " + string(v.SpacesCount))
		fmt.Println("Address: " + v.Address)
		fmt.Println("Description: " + v.Description)
	}
}

package objects

import (
	"DeathStar/models/tools"
	"fmt"
	"strconv"
)

type City struct {
	cityId    string
	name      string
	parkings	[]Parking
	districts	[]District
	spacesCount       int
	freeSpacesCount   int
	occSpacesCount    int
}

func NewCity(name string) *City {
	newCity := City{tools.GenerateId("CI"), name, make([]Parking, 0), make([]District, 0), 0, 0, 0}
	return &newCity
}

func (c *City) GetCityId() string {
	return c.cityId
}

func (c *City) GetName() string {
	return c.name
}

func (c *City) GetDistricts() *[]District {
	return &c.districts
}

func (c *City) GetDistrict(districtName string) *District {
	for i, _ := range c.districts {
		if c.districts[i].name == districtName {
			return  &c.districts[i]
		}
	}
	return nil
}

func (c *City) GetSpacesCount() int {
	return c.spacesCount
}

func (c *City) GetFreeSpacesCount() int {
	return c.freeSpacesCount
}

func (c *City) GetOccSpacesCount() int {
	return c.occSpacesCount
}

func (c *City) GetParkings() *[]Parking {
	return &c.parkings
}

func (c *City) GetParking(parkingName string) *Parking {
	for i, _ := range c.parkings {
		if c.parkings[i].name == parkingName {
			return &c.parkings[i]
		}
	}
	return nil
}

func (c *City) SetName(newName string) {
	c.name = newName
}

func (c *City) SetSpacesCount(newSpacesCount int) {
	c.spacesCount = newSpacesCount
}

func (c *City) SetFreeSpacesCount(newFreeSpacesCount int) {
	c.freeSpacesCount = newFreeSpacesCount
}

func (c *City) SetOccSpacesCount(newOccSpacesCount int) {
	c.occSpacesCount = newOccSpacesCount
}

func (c City) String() []string {
	var s = make([]string, 0)
	for _, p := range c.parkings {
		//s = append(s, "IDParking: " + p.parkingId)
		s = append(s, "NomParc: " + p.name)
		s = append(s,"Quartier: " + p.district.name)
		s = append(s, "Owner: " + p.owner.name)
		s = append(s, "StatusParc: " + p.status)
		s = append(s, "TotalParc: " +  strconv.Itoa(p.spacesCount))
		s = append(s, "PresentsParc: " + strconv.Itoa(p.occSpacesCount))
		s = append(s, "LibreParc: " + strconv.Itoa(p.freeSpacesCount))
		s = append(s, "")
	}
	return s
}

func (c City) PrintCity() {
	fmt.Printf("\nID: %#v\n", c.cityId)
	fmt.Printf("Name: %#v\n", c.name)

	fmt.Printf("Total: %#v\n", c.spacesCount)
	fmt.Printf("TotalPresent: %#v\n", c.occSpacesCount)
	fmt.Printf("TotalLibre: %#v\n\n", c.freeSpacesCount)

	/*
	fmt.Printf("Liste des Quartiers\n")
	for _, d := range c.districts {
		fmt.Printf("\n\tIDQuartier: %#v\n", d.districtId)
		fmt.Printf("\tNomQuartier: %#v\n", d.name)
		fmt.Printf("\tVille: %#v\n", d.city.name)
		fmt.Printf("\tTotalQuartier: %#v\n", d.spacesCount)
		fmt.Printf("\tTotalPresentsQuartier: %#v\n", d.occSpacesCount)
		fmt.Printf("\tTotalLibreQuartier: %#v\n", d.freeSpacesCount)
*/
	fmt.Printf("\n\tListe des Parkings\n")
	for _, p := range c.parkings {
		fmt.Printf("\n\t\tIDParking: %#v\n", p.parkingId)
		fmt.Printf("\t\tNomParc: %#v\n", p.name)
		fmt.Printf("\t\tQuartier: %#v\n", p.district.name)
		fmt.Printf("\t\tOwner: %#v\n", p.owner.name)
		fmt.Printf("\t\tStatusParc: %#v\n", p.status)
		fmt.Printf("\t\tTotalParc: %#v\n", p.spacesCount)
		fmt.Printf("\t\tPresentsParc: %#v\n", p.occSpacesCount)
		fmt.Printf("\t\tLibreParc: %#v\n", p.freeSpacesCount)
	}
}

func (c *City) AddDistrict(newDistrict *District) {
	c.districts = append(c.districts, *newDistrict)
}

func (c *City) DelDistrict(district District) {
	for i, v := range c.districts {
		if v.GetDistrictId() == district.GetDistrictId() {
			c.districts = append(c.districts[:i], c.districts[i+1:]...)
		}
	}
}

func (c *City) AddParking(newParking *Parking) {
	c.parkings = append(c.parkings, *newParking)
}

func (c *City) DelParking(parking Parking) {
	for i, v := range c.parkings {
		if v.GetParkingId() == parking.GetParkingId() {
			c.parkings = append(c.parkings[:i], c.parkings[i+1:]...)
		}
	}
}

func (c *City) UpdateCity() {
	newSpacesCount := 0
	newFreeSpacesCount := 0
	newOccSpacesCount := 0

	for i, _ := range c.parkings {
		//c.districts[i].UpdateDistrict()
		newSpacesCount += c.parkings[i].spacesCount
		newFreeSpacesCount += c.parkings[i].freeSpacesCount
		newOccSpacesCount += c.parkings[i].occSpacesCount
	}
	c.spacesCount = newSpacesCount
	c.freeSpacesCount = newFreeSpacesCount
	c.occSpacesCount = newOccSpacesCount
}
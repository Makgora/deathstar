package objects

import (
	"DeathStar/models/tools"
	"fmt"
)

type District struct {
	districtId string
	name       string
	city       *City
	parkings   []Parking
	spacesCount       int
	freeSpacesCount   int
	occSpacesCount    int
}

func NewDistrict(name string, city *City) *District {
	newDistrict := District{tools.GenerateId("DI"), name, city, make([]Parking, 0), 0, 0, 0}
	return &newDistrict
}

func (d *District) GetDistrictId() string {
	return d.districtId
}

func (d *District) GetName() string {
	return d.name
}

func (d *District) GetCity() *City {
	return d.city
}

func (d *District) GetParkings() *[]Parking {
	return &d.parkings
}

func (d *District) GetSpacesCount() int {
	return d.spacesCount
}

func (d *District) GetFreeSpacesCount() int {
	return d.freeSpacesCount
}

func (d *District) GetOccSpacesCount() int {
	return d.occSpacesCount
}

func (d *District) SetName(newName string) {
	d.name = newName
}

func (d *District) SetCity(newCity *City) {
	d.city = newCity
}

func (d *District) SetParkings(newParkingSlice *[]Parking) {
	d.parkings = *newParkingSlice
}

func (d *District) SetSpacesCount(newSpacesCount int) {
	d.spacesCount = newSpacesCount
}

func (d *District) SetFreeSpacesCount(newFreeSpacesCount int) {
	d.freeSpacesCount = newFreeSpacesCount
}

func (d *District) SetOccSpacesCount(newOccSpacesCount int) {
	d.occSpacesCount = newOccSpacesCount
}

func (d District) String() string {
	return fmt.Sprintf("[ID]: '%s' | [Name]: '%s' | [City]: '%s'| [SpacesCount]: '%d' |" +
		" [FreeSpacesCount]: '%d' | [OccSpacesCount]: '%d' ", d.districtId, d.name, d.city.GetName(),
			d.spacesCount, d.freeSpacesCount, d.occSpacesCount)
}

func (d *District) AddParking(newParking *Parking) {
	d.parkings = append(d.parkings, *newParking)
}

func (d *District) DelParking(parking Parking) {
	for i, v := range d.parkings {
		if v.GetParkingId() == parking.GetParkingId() {
			d.parkings = append(d.parkings[:i], d.parkings[i+1:]...)
		}
	}
}

func (d *District) UpdateDistrict() {
	newSpacesCount := 0
	newFreeSpacesCount := 0
	newOccSpacesCount := 0

	for _, v := range d.parkings {
		newSpacesCount += v.spacesCount
		newFreeSpacesCount += v.freeSpacesCount
		newOccSpacesCount += v.occSpacesCount
	}

	d.spacesCount = newSpacesCount
	d.freeSpacesCount = newFreeSpacesCount
	d.occSpacesCount = newOccSpacesCount
}
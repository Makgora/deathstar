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
}

func NewDistrict(name string, city *City) *District {
	newDistrict := District{tools.GenerateId("DI"), name, city, make([]Parking, 0)}
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

func (d *District) SetName(newName string) {
	d.name = newName
}

func (d *District) SetCity(newCity *City) {
	d.city = newCity
}

func (d District) String() string {
	return fmt.Sprintf("[ID]: '%s' | [Name]: '%s' | [City]: '%s'", d.districtId, d.name, d.city.GetName())
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

package objects

import (
	"fmt"
)

type District struct {
	name            string
	city            *City
	spacesCount     int
	freeSpacesCount int
	occSpacesCount  int
}

func NewDistrict(name string, city *City) *District {
	newDistrict := District{name, city, 0, 0, 0}
	GetLiveDB().GetCity(city.name).AddDistrict(&newDistrict)
	return &newDistrict
}

func (d *District) GetName() string {
	return d.name
}

func (d *District) GetCity() *City {
	return d.city
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
	return fmt.Sprintf("[Name]: '%s' | [City]: '%s'| [SpacesCount]: '%d' |" +
		" [FreeSpacesCount]: '%d' | [OccSpacesCount]: '%d' ", d.name, d.city.GetName(),
			d.spacesCount, d.freeSpacesCount, d.occSpacesCount)
}

func (d *District) UpdateDistrict() {
	/*newSpacesCount := 0
	newFreeSpacesCount := 0
	newOccSpacesCount := 0

	for _, v := range d.parkings {
		newSpacesCount += v.spacesCount
		newFreeSpacesCount += v.freeSpacesCount
		newOccSpacesCount += v.occSpacesCount
	}

	d.spacesCount = newSpacesCount
	d.freeSpacesCount = newFreeSpacesCount
	d.occSpacesCount = newOccSpacesCount*/
}
package objects

import (
	"fmt"
)

type Parking struct {
	id                string
	name              string
	owner             *Owner
	district          *District
	city			  *City
	status            string
	spacesCount       int
	freeSpacesCount   int
	occSpacesCount    int
	subOccSpacesCount int
	dayOccSpacesCount int
}

func NewParking(id string, name string, owner *Owner, district *District, city *City, spacesCount int) *Parking {
	newParking := Parking{id, name, owner, district, city, "", spacesCount, 0, 0, 0, 0}
	GetLiveDB().GetCity(city.name).AddParking(&newParking)
	return &newParking
}

func (p *Parking) GetId() string {
	return p.id
}

func (p *Parking) GetName() string {
	return p.name
}

func (p *Parking) GetOwner() *Owner {
	return p.owner
}

func (p *Parking) GetDistrict() *District {
	return p.district
}

func (p *Parking) GetStatus() string {
	return p.status
}

func (p *Parking) GetSpacesCount() int {
	return p.spacesCount
}

func (p *Parking) GetFreeSpacesCount() int {
	return p.freeSpacesCount
}

func (p *Parking) GetOccSpacesCount() int {
	return p.occSpacesCount
}

func (p *Parking) GetSubOccSpacesCount() int {
	return p.subOccSpacesCount
}

func (p *Parking) GetDayOccSpacesCount() int {
	return p.dayOccSpacesCount
}

func (p *Parking) SetName(newName string) {
	p.name = newName
}

func (p *Parking) SetOwner(newOwner *Owner) {
	p.owner = newOwner
}

func (p *Parking) SetDistrict(newDistrict *District) {
	p.district = newDistrict
}

func (p *Parking) SetStatus(newStatus string) {
	p.status = newStatus
}

func (p *Parking) SetSpacesCount(newSpacesCount int) {
	p.spacesCount = newSpacesCount
}

func (p *Parking) SetFreeSpacesCount(newFreeSpacesCount int) {
	p.freeSpacesCount = newFreeSpacesCount
}

func (p *Parking) SetOccSpacesCount(newOccSpacesCount int) {
	p.occSpacesCount = newOccSpacesCount
}

func (p *Parking) SetSubOccSpacesCount(newSubOccSpacesCount int) {
	p.subOccSpacesCount = newSubOccSpacesCount
}

func (p *Parking) SetDayOccSpacesCount(newDayOccSpacesCount int) {
	p.dayOccSpacesCount = newDayOccSpacesCount
}

func (p Parking) String() string {
	return fmt.Sprintf("[ID]: '%s' | [Name]: '%s' | [Owner]: '%s' "+
		"| [District]: '%s'| [Status]: %s | [SpacesCount]: '%d' | [FreeSpacesCount]: '%d' "+
		"| [OccSpacesCount]: '%d' | [SubOccSpacesCount]: '%d' | [DayOccSpacesCount]: '%d'",
		p.id, p.name, p.owner.GetName(), p.district.GetName(), p.status, p.spacesCount,
		p.freeSpacesCount, p.occSpacesCount, p.subOccSpacesCount, p.dayOccSpacesCount)
}

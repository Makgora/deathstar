package objects

import (
	"DeathStar/models/tools"
	"fmt"
)

type Parking struct {
	parkingId         string
	name              string
	owner             *Owner
	district          *District
	spacesCount       int
	freeSpacesCount   int
	occSpacesCount    int
	subOccSpacesCount int
	dayOccSpacesCount int
}

func NewParking(name string, owner *Owner, district *District, spacesCount int) *Parking {
	newParking := Parking{tools.GenerateId("PA"), name, owner, district, spacesCount, 0, 0, 0, 0}
	return &newParking
}

func (p *Parking) GetParkingId() string {
	return p.parkingId
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
	return fmt.Sprintf("[ID]: '%s' | [Name]: '%s' | [Owner]: '%s' " +
		"| [District]: '%s'| [SpacesCount]: '%d' | [FreeSpacesCount]: '%d' " +
			"| [OccSpacesCount]: '%d' | [SubOccSpacesCount]: '%d' | [DayOccSpacesCount]: '%d'",
		p.parkingId, p.name, p.owner.GetName(), p.district.GetName(), p.spacesCount,
			p.freeSpacesCount, p.occSpacesCount, p.subOccSpacesCount, p.dayOccSpacesCount)
}

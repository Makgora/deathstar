package objects

import (
	"DeathStar/models/tools"
	"fmt"
)

type Owner struct {
	ownerID  string
	name     string
	parkings []Parking
}

func NewOwner(name string) *Owner {
	newOwner := Owner{tools.GenerateId("OW"), name, make([]Parking, 0)}
	return &newOwner
}

func (o *Owner) GetOwnerId() string {
	return o.ownerID
}

func (o *Owner) GetName() string {
	return o.name
}

func (o *Owner) GetParkings() *[]Parking {
	return &o.parkings
}

func (o *Owner) SetName(newName string) {
	o.name = newName
}

func (o Owner) String() string {
	return fmt.Sprintf("[ID]: '%s' | [Name]: '%s'", o.ownerID, o.name)
}

func (o *Owner) AddParking(newParking *Parking) {
	o.parkings = append(o.parkings, *newParking)
}

func (o *Owner) DelParking(parking Parking) {
	for i, v := range o.parkings {
		if v.GetParkingId() == parking.GetParkingId() {
			o.parkings = append(o.parkings[:i], o.parkings[i+1:]...)
		}
	}
}

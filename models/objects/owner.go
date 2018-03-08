package objects

import (
	"fmt"
)

type Owner struct {
	name     string
}

func NewOwner(name string) *Owner {
	newOwner := Owner{name}
	GetLiveDB().AddOwner(&newOwner)
	return &newOwner
}

func (o *Owner) GetName() string {
	return o.name
}

func (o *Owner) SetName(newName string) {
	o.name = newName
}

func (o Owner) String() string {
	return fmt.Sprintf("[Name]: '%s'", o.name)
}


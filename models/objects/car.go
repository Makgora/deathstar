package objects

import (
	"DeathStar/models/tools"
	"fmt"
)

type Car struct {
	carId string
	name  string
	user  *User
}

func NewCar(name string, user *User) *Car {
	newCar := Car{tools.GenerateId("CA"), name, user}
	return &newCar
}

func (c *Car) GetCarId() string {
	return c.carId
}

func (c *Car) GetName() string {
	return c.name
}

func (c *Car) GetUser() *User {
	return c.user
}

func (c *Car) SetName(newName string) {
	c.name = newName
}

func (c *Car) SetUser(newUser *User) {
	c.user = newUser
}

func (c Car) String() string {
	return fmt.Sprintf("[ID]: '%s' | [Name]: '%s' | [User]: '%s'", c.carId, c.name, c.user.GetName())
}

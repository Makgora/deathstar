package objects

import (
	"DeathStar/models/tools"
	"fmt"
)

type User struct {
	userId string
	name   string
	cars   []Car
}

func NewUser(name string) *User {
	newUser := User{tools.GenerateId("US"), name, make([]Car, 0)}
	return &newUser
}

func (u *User) GetUserId() string {
	return u.userId
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetCars() *[]Car {
	return &u.cars
}

func (u *User) SetName(newName string) {
	u.name = newName
}

func (u User) String() string {
	return fmt.Sprintf("[ID]: '%s' | [Name]: '%s'", u.userId, u.name)
}

func (u *User) AddCar(newCar *Car) {
	u.cars = append(u.cars, *newCar)
}

func (u *User) DelCar(car Car) {
	for i, v := range u.cars {
		if v.GetCarId() == car.GetCarId() {
			u.cars = append(u.cars[:i], u.cars[i+1:]...)
		}
	}
}

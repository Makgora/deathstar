package users_tests

import (
	"fmt"
	"testing"
	. "DeathStar/models/users"
	. "DeathStar/models/cars"
)

func TestNewUser(t *testing.T) {
	name := "testUser"

	user := NewUser(name)
	if user.GetName() != name {
		t.Error("Wrong name ! Expected ", name, " got ", user.GetName())
	}
}

func TestSetters(t *testing.T) {
	user := NewUser("")

	t.Run("SetName", func(t *testing.T) {
		newNameTest := "newNameTest"
		user.SetName(newNameTest)
		if user.GetName() != newNameTest {
			t.Error("Wrong name ! Expected ", newNameTest, " got ", user.GetName())
		}
	})
}

func TestAddCar(t *testing.T) {
	user := NewUser("")
	newCar := NewCar("", user)

	if len(user.GetCars()) != 0 {
		t.Error("Wrong cars number ! Expected 0 got ", len(user.GetCars()))
	}
	user.AddCar(newCar)
	if len(user.GetCars()) != 1 {
		t.Error("Wrong cars number ! Expected 1 got ", len(user.GetCars()))
	}
}

func TestDelCar(t *testing.T) {
	user := NewUser("")
	newCar := NewCar("", user)
	user.AddCar(newCar)
	if len(user.GetCars()) != 1 {
		t.Error("Wrong cars number ! Expected 1 got ", len(user.GetCars()))
	}
	user.DelCar(newCar)
	if len(user.GetCars()) != 0 {
		t.Error("Wrong cars number ! Expected 0 got ", len(user.GetCars()))
	}
}

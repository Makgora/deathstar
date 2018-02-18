package cars_tests_test

import (
	. "DeathStar/models/cars"
	. "DeathStar/models/users"
	"testing"
)

func TestNewCar(t *testing.T) {
	name := "nameTest"
	user := NewUser("userTest")

	car := NewCar(name, user)

	if car.GetName() != name {
		t.Error("Wrong name ! Expected ", name, " got ", car.GetName())
	} else if car.GetUser().GetName() != user.GetName() {
		t.Error("Wrong user ! Expected ", user.GetName(), " got ", car.GetUser().GetName())
	}
}

func TestSetters(t *testing.T) {
	user := NewUser("userTest")

	car := NewCar("testName", user)

	t.Run("SetName", func(t *testing.T) {
		car.SetName("newName")
		if car.GetName() != "SetName" {
			t.Error("Wrong name ! Expected newName got ", car.GetName())
		}
	})
	t.Run("SetUser", func(t *testing.T) {
		car.SetUser(NewUser("newUserTest"))
		if car.GetUser().GetName() != newUserName {
			t.Error("Wrong User ! Expected newUserTest got ", car.GetUser().GetName())
		}
	})
}

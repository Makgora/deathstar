package owners_tests

import (
	. "DeathStar/models/cities"
	. "DeathStar/models/countries"
	. "DeathStar/models/districts"
	. "DeathStar/models/owners"
	. "DeathStar/models/parkings"
	"testing"
)

func TestNewOwner(t *testing.T) {
	name := "nameTest"

	owner := NewOwner(name)
	if owner.GetName() != name {
		t.Error("Wrong name ! Expected ", name, " got ", owner.GetName())
	}
}

func TestSetters(t *testing.T) {
	owner := NewOwner("nameTest")
	t.Run("SetName", func(t *testing.T) {
		newNameTest := "newNameTest"
		owner.SetName(newNameTest)
		if owner.GetName() != newNameTest {
			t.Error("Wrong name ! Expected ", newNameTest, " got ", owner.GetName())
		}
	})
}

func TestAddParking(t *testing.T) {
	owner := NewOwner("")
	if len(owner.GetParkings()) != 0 {
		t.Error("Wrong parkings number ! Expected 0 got ", len(owner.GetParkings()))
	}
	country := NewCountry("")
	city := NewCity("", country)
	district := NewDistrict("", city)
	newParking := NewParking("", owner, district, 0)
	owner.AddParking(newParking)
	if len(owner.GetParkings()) != 1 {
		t.Error("Wrong parkings number ! Expected 1 got ", len(owner.GetParkings()))
	}
}

func TestDelParking(t *testing.T) {
	owner := NewOwner("")
	country := NewCountry("")
	city := NewCity("", country)
	district := NewDistrict("", city)
	newParking := NewParking("", owner, district, 0)

	owner.AddParking(newParking)
	if len(owner.GetParkings()) != 1 {
		t.Error("Wrong parkings number ! Expected 1 got ", len(owner.GetParkings()))
	}
	owner.DelParking(newParking)
	if len(owner.GetParkings()) != 0 {
		t.Error("Wrong parkings number ! Expected 0 got ", len(owner.GetParkings()))
	}
}

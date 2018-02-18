package districts_tests

import (
	. "DeathStar/models/cars"
	. "DeathStar/models/cities"
	. "DeathStar/models/countries"
	. "DeathStar/models/districts"
	. "DeathStar/models/owners"
	. "DeathStar/models/parkings"
	"fmt"
	"testing"
)

func TestNewDistrict(t *testing.T) {
	name := "nameTest"
	country := NewCountry("countryTest")
	city := NewCity("cityTest", country)

	district := NewDistrict(name, city)

	fmt.Sprintf(district.GetDistrictId())
	if district.GetName() != name {
		t.Error("Wrong name ! Expected ", name, " got ", district.GetName())
	} else if district.GetCity().GetName() != city.GetName() {
		t.Error("Wrong city ! Expected ", city.GetName(), " got ", district.GetCity().GetName())
	}
}

func TestSetters(t *testing.T) {
	country := NewCountry("")
	city := NewCity("", country)
	district := NewDistrict("nameTest", city)

	t.Run("SetName", func(t *testing.T) {
		newNameTest := "newNameTest"
		district.SetName(newNameTest)
		if district.GetName() != newNameTest {
			t.Error("Wrong name ! Expected ", newNameTest, " got ", district.GetName())
		}
	})
	t.Run("SetCity", func(t *testing.T) {
		newCity := NewCity("newCityName", country)
		district.SetCity(&newCity)
		if district.GetCity().GetName() != newCity.GetName() {
			t.Error("Wrong city ! Expected ", newCity.GetName(), " got ", district.GetCity().GetName())
		}
	})
}

func TestAddParking(t *testing.T) {
	country := NewCountry("")
	city := NewCity("", country)
	district := NewDistrict("nameTest", city)

	if len(district.GetParkings()) != 0 {
		t.Error("Wrong parkings number ! Expected 0 got ", len(district.GetParkings()))
	}
	user := NewUser("")
	car := NewCar("", user)
	owner := NewOwner("", car)
	newParking := NewParking("", owner, district, 0)
	district.AddParking(newParking)
	if len(district.GetParkings()) != 1 {
		t.Error("Wrong parkings number ! Expected 1 got ", len(district.GetParkings()))
	}
}

func TestDelParking(t *testing.T) {
	country := NewCountry("")
	city := NewCity("", country)
	user := NewUser("")
	car := NewCar("", user)
	owner := NewOwner("", car)
	district := NewDistrict("nameTest", city)

	newParking := NewParking("", owner, district, 0)
	district.AddParking(newParking)
	if len(district.GetParkings()) != 1 {
		t.Error("Wrong parkings number ! Expected 1 got ", len(district.GetParkings()))
	}
	district.DelParking(newParking)
	if len(district.GetParkings()) != 0 {
		t.Error("Wrong parkings number ! Expected 0 got ", len(district.GetParkings()))
	}
}

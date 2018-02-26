package objects_tests

import (
	. "DeathStar/models/objects"
	"testing"
)

func TestNewDistrict(t *testing.T) {
	name := "nameTest"
	country := NewCountry("countryTest")
	city := NewCity("cityTest", country)

	district := NewDistrict(name, city)

	if district.GetName() != name {
		t.Error("Wrong name ! Expected ", name, " got ", district.GetName())
	} else if district.GetCity().GetName() != city.GetName() {
		t.Error("Wrong city ! Expected ", city.GetName(), " got ", district.GetCity().GetName())
	}
}

func TestDistrictSetters(t *testing.T) {
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
		district.SetCity(newCity)
		if district.GetCity().GetName() != newCity.GetName() {
			t.Error("Wrong city ! Expected ", newCity.GetName(), " got ", district.GetCity().GetName())
		}
	})
}

func TestDistrictAddParking(t *testing.T) {
	country := NewCountry("")
	city := NewCity("", country)
	district := NewDistrict("nameTest", city)

	if len(*district.GetParkings()) != 0 {
		t.Error("Wrong parkings number ! Expected 0 got ", len(*district.GetParkings()))
	}
	owner := NewOwner("")
	newParking := NewParking("", owner, district, 0)
	district.AddParking(newParking)
	if len(*district.GetParkings()) != 1 {
		t.Error("Wrong parkings number ! Expected 1 got ", len(*district.GetParkings()))
	}
}

func TestDistrictDelParking(t *testing.T) {
	country := NewCountry("")
	city := NewCity("", country)
	owner := NewOwner("")
	district := NewDistrict("nameTest", city)

	newParking := NewParking("", owner, district, 0)
	district.AddParking(newParking)
	if len(*district.GetParkings()) != 1 {
		t.Error("Wrong parkings number ! Expected 1 got ", len(*district.GetParkings()))
	}
	district.DelParking(*newParking)
	if len(*district.GetParkings()) != 0 {
		t.Error("Wrong parkings number ! Expected 0 got ", len(*district.GetParkings()))
	}
}

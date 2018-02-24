package objects_tests

import (
	. "DeathStar/models/objects"
	"testing"
)

func TestNewCity(t *testing.T) {
	name := "cityTest"
	country := NewCountry("countryTest")

	city := NewCity(name, &country)

	if city.GetName() != name {
		t.Error("Wrong name ! Expected ", name, " got ", city.GetName())
	} else if city.GetCountry().GetName() != country.GetName() {
		t.Error("Wrong city ! Expected ", country.GetName(), " got ", city.GetCountry().GetName())
	}
}

func TestCitySetters(t *testing.T) {
	country := NewCountry("testCountry")
	city := NewCity("nameTest", &country)

	t.Run("SetName", func(t *testing.T) {
		newNameTest := "newNameTest"
		city.SetName(newNameTest)
		if city.GetName() != newNameTest {
			t.Error("Wrong name ! Expected ", newNameTest, " got ", city.GetName())
		}
	})
	t.Run("SetCountry", func(t *testing.T) {
		newCountry := NewCountry("newTestCountry")
		city.SetCountry(&newCountry)
		if city.GetCountry().GetName() != newCountry.GetName() {
			t.Error("Wrong country ! Expected ", newCountry, " got ", city.GetCountry().GetName())
		}
	})
}

func TestCityAddDistrict(t *testing.T) {
	country := NewCountry("")
	city := NewCity("", &country)

	if len(city.GetDistricts()) != 0 {
		t.Error("Wrong district number ! Expected 0 got ", len(city.GetDistricts()))
	}
	newDistrict := NewDistrict("", &city)
	city.AddDistrict(&newDistrict)
	if len(city.GetDistricts()) != 1 {
		t.Error("Wrong district number ! Expected 1 got ", len(city.GetDistricts()))
	}
}

func TestCityDelDistrict(t *testing.T) {
	country := NewCountry("")
	city := NewCity("", &country)

	newDistrict := NewDistrict("", &city)
	city.AddDistrict(&newDistrict)
	if len(city.GetDistricts()) != 1 {
		t.Error("Wrong district number ! Expected 1 got ", len(city.GetDistricts()))
	}
	city.DelDistrict(newDistrict)
	if len(city.GetDistricts()) != 0 {
		t.Error("Wrong district number ! Expected 0 got ", len(city.GetDistricts()))
	}
}

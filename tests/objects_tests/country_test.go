package objects_tests

import (
	. "DeathStar/models/objects"
	"testing"
)

func TestNewCountry(t *testing.T) {
	name := "nameTest"

	country := NewCountry(name)
	if country.GetName() != name {
		t.Error("Wrong name ! Expected ", name, " got ", country.GetName())
	}
}

func TestCountrySetters(t *testing.T) {
	country := NewCountry("")

	t.Run("SetName", func(t *testing.T) {
		newNameTest := "newNameTest"
		country.SetName(newNameTest)
		if country.GetName() != newNameTest {
			t.Error("Wrong name ! Expected ", newNameTest, " got ", country.GetName())
		}
	})
}

func TestCountryAddCity(t *testing.T) {
	country := NewCountry("")

	if len(*country.GetCities()) != 0 {
		t.Error("Wrong cities number ! Expected 0 got ", len(*country.GetCities()))
	}
	newCity := NewCity("", country)
	country.AddCity(newCity)
	if len(*country.GetCities()) != 1 {
		t.Error("Wrong cities number ! Expected 1 got ", len(*country.GetCities()))
	}
}

func TestCountryDelCity(t *testing.T) {
	country := NewCountry("")
	newCity := NewCity("", country)

	country.AddCity(newCity)
	if len(*country.GetCities()) != 1 {
		t.Error("Wrong cities number ! Expected 1 got ", len(*country.GetCities()))
	}
	country.DelCity(*newCity)
	if len(*country.GetCities()) != 0 {
		t.Error("Wrong cities number ! Expected 0 got ", len(*country.GetCities()))
	}
}

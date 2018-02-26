package objects

import (
	"DeathStar/models/tools"
	"fmt"
)

type Country struct {
	countryId string
	name      string
	cities    []City
}

func NewCountry(name string) *Country {
	newCountry := Country{tools.GenerateId("CO"), name, make([]City, 0)}
	return &newCountry
}

func (c *Country) GetCountryId() string {
	return c.countryId
}

func (c *Country) GetName() string {
	return c.name
}

func (c *Country) GetCities() *[]City {
	return &c.cities
}

func (c *Country) SetName(newName string) {
	c.name = newName
}

func (c Country) String() string {
	return fmt.Sprintf("[ID]: '%s' | [Name]: '%s'", c.countryId, c.name)
}

func (c *Country) AddCity(newCity *City) {
	c.cities = append(c.cities, *newCity)
}

func (c *Country) DelCity(city City) {
	for i, v := range c.cities {
		if v.GetCityId() == city.GetCityId() {
			c.cities = append(c.cities[:i], c.cities[i+1:]...)
		}
	}
}

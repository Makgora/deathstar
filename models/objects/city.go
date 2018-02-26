package objects

import (
	"DeathStar/models/tools"
	"fmt"
)

type City struct {
	cityId    string
	name      string
	country   *Country
	districts []District
}

func NewCity(name string, country *Country) *City {
	newCity := City{tools.GenerateId("CI"), name, country, make([]District, 0)}
	return &newCity
}

func (c *City) GetCityId() string {
	return c.cityId
}

func (c *City) GetName() string {
	return c.name
}

func (c *City) GetCountry() *Country {
	return c.country
}

func (c *City) GetDistricts() *[]District {
	return &c.districts
}

func (c *City) SetName(newName string) {
	c.name = newName
}

func (c *City) SetCountry(newCountry *Country) {
	c.country = newCountry
}

func (c City) String() string {
	return fmt.Sprintf("[ID]: '%s' | [Name]: '%s' | [Country]: '%s'", c.cityId, c.name, c.country.GetName())
}

func (c *City) AddDistrict(newDistrict *District) {
	c.districts = append(c.districts, *newDistrict)
}

func (c *City) DelDistrict(district District) {
	for i, v := range c.districts {
		if v.GetDistrictId() == district.GetDistrictId() {
			c.districts = append(c.districts[:i], c.districts[i+1:]...)
		}
	}
}

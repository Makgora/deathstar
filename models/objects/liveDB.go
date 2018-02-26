package objects

import "github.com/astaxie/beego"

type LiveDB struct {
	cars		[]Car
	cities		[]City
	countries	[]Country
	districts	[]District
	owners		[]Owner
	parkings	[]Parking
	users		[]User
}

func NewLiveDB() LiveDB {
	newLiveDB := LiveDB{make([]Car, 0), make([]City, 0), make([]Country, 0),
	make([]District, 0), make([]Owner, 0), make([]Parking, 0),
	make([]User, 0)}
	return newLiveDB
}

func (l *LiveDB) GetCars() *[]Car {
	return &l.cars
}

func (l *LiveDB) GetCities() *[]City {
	return &l.cities
}

func (l *LiveDB) GetCountries() *[]Country {
	return &l.countries
}

func (l *LiveDB) GetOwners() *[]Owner {
	return &l.owners
}

func (l *LiveDB) GetParkings() *[]Parking {
	return &l.parkings
}

func (l *LiveDB) GetUsers() *[]User {
	return &l.users
}

func (l *LiveDB) AddCar(newCar *Car) {
	l.cars = append(l.cars, *newCar)
}

func (l *LiveDB) AddCity(newCity *City) {
	l.cities = append(l.cities, *newCity)
}

func (l *LiveDB) AddCountry(newCountry *Country) {
	l.countries = append(l.countries, *newCountry)
}

func (l *LiveDB) AddOwner(newOwner *Owner) {
	l.owners = append(l.owners, *newOwner)
}

func (l *LiveDB) AddParking(newParking *Parking) {
	l.parkings = append(l.parkings, *newParking)
}

func (l *LiveDB) AddUser(newUser *User) {
	l.users = append(l.users, *newUser)
}

func (l *LiveDB) DelCar(car Car) {
	for i, v := range l.cars {
		if v.GetCarId() == car.GetCarId() {
			l.cars = append(l.cars[:i], l.cars[i+1:]...)
		}
	}
}

func (l *LiveDB) DelCity(city City) {
	for i, v := range l.cities {
		if v.GetCityId() == city.GetCityId() {
			l.cities = append(l.cities[:i], l.cities[i+1:]...)
		}
	}
}

func (l *LiveDB) DelCountry(country Country) {
	for i, v := range l.countries {
		if v.GetCountryId() == country.GetCountryId() {
			l.countries = append(l.countries[:i], l.countries[i+1:]...)
		}
	}
}

func (l *LiveDB) DelOwner(owner Owner) {
	for i, v := range l.owners {
		if v.GetOwnerId() == owner.GetOwnerId() {
			l.owners = append(l.owners[:i], l.owners[i+1:]...)
		}
	}
}

func (l *LiveDB) DelParking(parking Parking) {
	for i, v := range l.parkings {
		if v.GetParkingId() == parking.GetParkingId() {
			l.parkings = append(l.parkings[:i], l.parkings[i+1:]...)
		}
	}
}

func (l *LiveDB) DelUser(user User) {
	for i, v := range l.users {
		if v.GetUserId() == user.GetUserId() {
			l.users = append(l.users[:i], l.users[i+1:]...)
		}
	}
}
package objects

import (
	"sync"
)

type LiveDB struct {
	cities []City
	owners []Owner
	users  []User
}

var liveDB *LiveDB
var once sync.Once

func newLiveDB() *LiveDB {
	newLiveDB := LiveDB{make([]City, 0), make([]Owner, 0), make([]User, 0)}
	return &newLiveDB
}

func GetLiveDB() *LiveDB {
	once.Do(func() {
		liveDB = newLiveDB()
	})
	return liveDB
}

func (l *LiveDB) GetCities() *[]City {
	return &l.cities
}

//TODO Optimize access
func (l *LiveDB) GetCity(cityName string) *City {
	for i, _ := range l.cities {
		if l.cities[i].name == cityName {
			return &l.cities[i]
		}
	}
	return nil
}

func (l *LiveDB) GetOwners() *[]Owner {
	return &l.owners
}

//TODO Optimize access
func (l *LiveDB) GetOwner(ownerName string) *Owner {
	for i, _ := range l.owners {
		if l.owners[i].name == ownerName {
			return &l.owners[i]
		}
	}
	return nil
}

func (l *LiveDB) GetUsers() *[]User {
	return &l.users
}

//TODO Optimize access
func (l *LiveDB) GetUser(userName string) *User {
	for i, _ := range l.users {
		if l.users[i].name == userName {
			return &l.users[i]
		}
	}
	return nil
}

func (l *LiveDB) AddCity(newCity *City) {
	l.cities = append(l.cities, *newCity)
}

func (l *LiveDB) AddOwner(newOwner *Owner) {
	l.owners = append(l.owners, *newOwner)
}

func (l *LiveDB) AddUser(newUser *User) {
	l.users = append(l.users, *newUser)
}

func (l *LiveDB) DelCity(city City) {
	for i, v := range l.cities {
		if v.name == city.name {
			l.cities = append(l.cities[:i], l.cities[i+1:]...)
		}
	}
}

func (l *LiveDB) DelOwner(owner Owner) {
	for i, v := range l.owners {
		if v.name == owner.name {
			l.owners = append(l.owners[:i], l.owners[i+1:]...)
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

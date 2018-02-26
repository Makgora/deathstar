package objects

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

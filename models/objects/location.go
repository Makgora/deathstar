package objects

type Location struct {
	lat		float32
	lng		float32
}

func NewLocation(lat float32, lng float32) Location {
	newLocation := Location{lat, lng}
	return newLocation
}

func (l Location) GetLat() float32 {
	return l.lat
}

func (l Location) GetLng() float32 {
	return l.lng
}
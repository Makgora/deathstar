package monaco

import (
	"DeathStar/models/objects"
	. "DeathStar/models/cities/monaco/MonacoParking"
)

func InitMonaco() {
	objects.GetLiveDB().AddCity(objects.NewCity("Monaco"))

	// Init parkings
	InitMonacoParkings()
}

func UpdateMonaco() {
	if objects.GetLiveDB().GetCity("Monaco") == nil {
		panic("Monaco is not initialized!")
	}
	// Update parkings
	UpdateMonacoParkings()

	// Update Monaco
	objects.GetLiveDB().GetCity("Monaco").UpdateCity()
}

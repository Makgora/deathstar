package monaco

import (
	"DeathStar/models/objects"
	. "DeathStar/models/cities/monaco/MonacoParking"
	"DeathStar/models/tools"
)

func InitMonaco() {
	objects.NewCity("Monaco")
	monaco := objects.GetLiveDB().GetCity("Monaco")

	// Init districts
	districts := tools.ParseDistrictsXml("./models/cities/monaco/districts.xml")

	for _, v := range districts.Districts {
		objects.NewDistrict(v.Name, monaco)
	}

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


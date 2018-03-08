package MonacoParking

import (
	"DeathStar/models/objects"
	"DeathStar/models/tools"
)

func InitMonacoParkings() {
	monacoParking := objects.NewOwner("MonacoParking")
	monaco := objects.GetLiveDB().GetCity("Monaco")

	parkings := tools.ParseParkingsXml("./models/cities/monaco/MonacoParking/parkings.xml")

	// Add parkings
	for _, d := range parkings.Parkings {
		objects.NewParking(d.Id, d.Name, monacoParking, monaco.GetDistrict(d.District), monaco, d.SpacesCount)
	}
}

func UpdateMonacoParkings() {
	monaco := objects.GetLiveDB().GetCity("Monaco")

	// Parse XML
	s := dlParseXml()

	// Update parkings values
	for _, v := range s.Quartier {
		for _, p := range v.Parc {
			parkingToUpdate := monaco.GetParking(p.NumParc)

			parkingToUpdate.SetFreeSpacesCount(p.PlacesLibresParc)
			parkingToUpdate.SetOccSpacesCount(p.Presents.PresentsAbonnes + p.Presents.PresentsHoraires)
			parkingToUpdate.SetSpacesCount(parkingToUpdate.GetOccSpacesCount() + parkingToUpdate.GetFreeSpacesCount())
			parkingToUpdate.SetStatus(p.StatusParc)
		}
	}
}
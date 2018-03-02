package MonacoParking

import (
	"DeathStar/models/objects"
)

func InitMonacoParkings() {
	monaco := objects.GetLiveDB().GetCity("Monaco")

	// Parse XML
	s := parseXml()

	// Create new owner
	monacoParking := objects.NewOwner("MonacoParking")
	objects.GetLiveDB().AddOwner(monacoParking)

	// Add districts and parkings TODO owner and district reference
	for _, d := range s.Quartier {
		newDistrict := objects.NewDistrict(d.NomQuartier, monaco)
		for _, p := range d.Parc {
			newParking := objects.NewParking(p.LibelleParc, monacoParking, newDistrict, 0)
			monaco.AddParking(newParking)
			monacoParking.AddParking(newParking)
			newDistrict.AddParking(newParking)
		}
		monaco.AddDistrict(newDistrict)
	}
}

func UpdateMonacoParkings() {
	monaco := objects.GetLiveDB().GetCity("Monaco")

	// Parse XML
	s := parseXml()

	// Update parkings values
	for _, v := range s.Quartier {
		for _, p := range v.Parc {
			parkingToUpdate := monaco.GetParking(p.LibelleParc)

			parkingToUpdate.SetSpacesCount(p.Presents.PresentsTotal + p.PlacesLibresParc)
			parkingToUpdate.SetFreeSpacesCount(p.PlacesLibresParc)
			parkingToUpdate.SetOccSpacesCount(p.Presents.PresentsTotal)
			parkingToUpdate.SetStatus(p.StatusParc)
		}
	}
}
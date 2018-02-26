package objects_tests

import (
	. "DeathStar/models/objects"
	"testing"
)

func TestNewParking(t *testing.T) {
	country := NewCountry("")
	city := NewCity("", country)

	name := "nameTest"
	owner := NewOwner("")
	district := NewDistrict("", city)
	spacesCount := 100

	parking := NewParking(name, owner, district, spacesCount)
	if parking.GetName() != name {
		t.Error("Wrong name ! Expected ", name, " got ", parking.GetName())
	}
	if parking.GetOwner().GetOwnerId() != owner.GetOwnerId() {
		t.Error("Wrong owner ! Expected ", owner.GetOwnerId(), " got ", parking.GetOwner().GetOwnerId())
	}
	if parking.GetDistrict().GetDistrictId() != district.GetDistrictId() {
		t.Error("Wrong district ! Expected ", district.GetDistrictId(), " got ", parking.GetDistrict().GetDistrictId())
	}
	if parking.GetSpacesCount() != spacesCount {
		t.Error("Wrong spacesCount ! Expected ", spacesCount, " got ", parking.GetSpacesCount())
	}
	if parking.GetFreeSpacesCount() != 0 {
		t.Error("Wrong freeSpacesCount ! Expected 0 got ", parking.GetFreeSpacesCount())
	}
	if parking.GetOccSpacesCount() != 0 {
		t.Error("Wrong occSpacesCount ! Expected 0 got ", parking.GetOccSpacesCount())
	}
	if parking.GetSubOccSpacesCount() != 0 {
		t.Error("Wrong subOccspacesCount ! Expected 0 got ", parking.GetSubOccSpacesCount())
	}
	if parking.GetDayOccSpacesCount() != 0 {
		t.Error("Wrong dayOccspacesCount ! Expected 0 got ", parking.GetDayOccSpacesCount())
	}
}

func TestParkingSetters(t *testing.T) {
	country := NewCountry("")
	city := NewCity("", country)

	name := "nameTest"
	owner := NewOwner("")
	district := NewDistrict("", city)
	spacesCount := 100

	parking := NewParking(name, owner, district, spacesCount)
	t.Run("SetName", func(t *testing.T) {
		newName := "newNameTest"
		parking.SetName(newName)
		if parking.GetName() != newName {
			t.Error("Wrong name ! Expected ", newName, " got ", parking.GetName())
		}
	})
	t.Run("SetOwner", func(t *testing.T) {
		newOwner := NewOwner("")
		parking.SetOwner(newOwner)
		if parking.GetOwner().GetOwnerId() != newOwner.GetOwnerId() {
			t.Error("Wrong owner ! Expected ", newOwner.GetOwnerId(), " got ", parking.GetOwner().GetOwnerId())
		}
	})
	t.Run("SetDistrict", func(t *testing.T) {
		newDistrict := NewDistrict("", city)
		parking.SetDistrict(newDistrict)
		if parking.GetDistrict().GetDistrictId() != newDistrict.GetDistrictId() {
			t.Error("Wrong district ! Expected ", newDistrict.GetDistrictId(), " got ", parking.GetDistrict().GetDistrictId())
		}
	})
	t.Run("SetSpacesCount", func(t *testing.T) {
		parking.SetSpacesCount(150)
		if parking.GetSpacesCount() != 150 {
			t.Error("Wrong spacesCount ! Expected 150 got ", parking.GetSpacesCount())
		}
	})
	t.Run("SetFreeSpacesCount", func(t *testing.T) {
		parking.SetFreeSpacesCount(10)
		if parking.GetFreeSpacesCount() != 10 {
			t.Error("Wrong freeSpacesCount ! Expected 10 got ", parking.GetFreeSpacesCount())
		}
	})
	t.Run("SetOccSpacesCount", func(t *testing.T) {
		parking.SetOccSpacesCount(10)
		if parking.GetOccSpacesCount() != 10 {
			t.Error("Wrong occSpacesCount ! Expected 10 got ", parking.GetOccSpacesCount())
		}
	})
	t.Run("SetSubOccSpacesCount", func(t *testing.T) {
		parking.SetSubOccSpacesCount(10)
		if parking.GetSubOccSpacesCount() != 10 {
			t.Error("Wrong subOccSpacesCount ! Expected 10 got ", parking.GetSubOccSpacesCount())
		}
	})
	t.Run("SetDayOccSpacesCount", func(t *testing.T) {
		parking.SetDayOccSpacesCount(10)
		if parking.GetDayOccSpacesCount() != 10 {
			t.Error("Wrong dayOccSpacesCount ! Expected 10 got ", parking.GetDayOccSpacesCount())
		}
	})
}

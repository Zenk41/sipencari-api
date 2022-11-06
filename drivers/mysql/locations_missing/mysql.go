package locationsmissing

import (
	locationsmissing "sipencari-api/businesses/locations_missing"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LocationMissingRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) locationsmissing.Repository {
	return &LocationMissingRepository{
		conn: conn,
	}
}

func (lcmr *LocationMissingRepository) GetAll() []locationsmissing.Domain {
	var rec []LocationMissing
	lcmr.conn.
		Preload("User").
		Preload("Missing").
		Find(&rec)
	locationMissingDomain := []locationsmissing.Domain{}
	for _, missing := range rec {
		locationMissingDomain = append(locationMissingDomain, missing.ToDomain())
	}
	return locationMissingDomain
}

func (lcmr *LocationMissingRepository) GetByID(idMissing, idLocation string) locationsmissing.Domain {
	locationMissing := LocationMissing{}
	lcmr.conn.
		Preload("User").
		Preload("Missing").
		First(&locationMissing, "missing_id=? AND id=?", idMissing, idLocation)
	return locationMissing.ToDomain()
}

func (lcmr *LocationMissingRepository) Create(idMissing string, locationMissingDomain *locationsmissing.Domain) locationsmissing.Domain {
	rec := FromDomain(locationMissingDomain)
	rec.ID = uuid.NewString()
	rec.MissingID = idMissing
	result := lcmr.conn.
		Preload("User").
		Preload("Missing").
		Create(&rec)
	result.Last(&rec)
	return rec.ToDomain()
}

func (lcmr *LocationMissingRepository) Update(idMissing, idLocation string, locationMissingDomain *locationsmissing.Domain) locationsmissing.Domain {
	var locationMissing locationsmissing.Domain = lcmr.GetByID(idMissing, idLocation)
	updatedLocation := FromDomain(&locationMissing)
	updatedLocation.Name = locationMissingDomain.Name
	updatedLocation.Lat = locationMissingDomain.Lat
	updatedLocation.Lng = locationMissingDomain.Lng
	lcmr.conn.Save(&updatedLocation)
	return updatedLocation.ToDomain()
}

func (lcmr *LocationMissingRepository) Delete(idMissing, idLocation string) bool {
	var locationMissing locationsmissing.Domain = lcmr.GetByID(idMissing, idLocation)

	deletedLocationMissing := FromDomain(&locationMissing)
	if result := lcmr.conn.Unscoped().Delete(&deletedLocationMissing); result.RowsAffected == 0 {
		return false
	}
	return true
}

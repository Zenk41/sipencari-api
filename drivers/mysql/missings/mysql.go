package missings

import (
	"sipencari-api/businesses/missings"
	"sipencari-api/drivers/mysql/hashtags"
	locationsmissing "sipencari-api/drivers/mysql/locations_missing"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type missingRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) missings.Repository {
	return &missingRepository{
		conn: conn,
	}
}

func (mr *missingRepository) GetAll() []missings.Domain {
	var rec []Missing
	mr.conn.
		Preload("User").
		Preload("Category").
		Preload("Hashtags").
		Preload("LocationMissing").
		Preload("LikeMissings.User").
		Preload("Comments.LocationComment").
		Preload("Comments.User").
		Preload("Comments.LikeComments").
		Find(&rec)
	missingDomain := []missings.Domain{}
	for _, missing := range rec {
		missingDomain = append(missingDomain, missing.ToDomain())
	}
	return missingDomain
}

func (mr *missingRepository) GetByID(id string) missings.Domain {
	missing := Missing{}
	mr.conn.
		Preload("User").
		Preload("Category").
		Preload("Hashtags").
		Preload("LocationMissing").
		Preload("LikeMissings.User").
		Preload("Comments.LocationComment").
		Preload("Comments.User").
		Preload("Comments.LikeComments").
		First(&missing, "id=?", id)
	return missing.ToDomain()
}

func (mr *missingRepository) Create(idUser string, missingDomain *missings.Domain) missings.Domain {
	rec := FromDomain(missingDomain)
	rec.ID = uuid.NewString()
	rec.UserID = idUser
	rec.LocationMissing.ID = uuid.NewString()
	result := mr.conn.
		Preload("User").
		Preload("Category").
		Preload("Hashtags").
		Preload("LocationMissing").
		Preload("LikeMissings.User").
		Preload("Comments.LocationComment").
		Preload("Comments.User").
		Preload("Comments.LikeComments").
		Create(&rec)
	result.Last(&rec)
	return rec.ToDomain()
}

func (mr *missingRepository) Update(idUser string, id string, missingDomain *missings.Domain) missings.Domain {
	var missing missings.Domain = mr.GetByID(id)
	updatedMissing := FromDomain(&missing)
	updatedMissing.Title = missingDomain.Title
	updatedMissing.Content = missingDomain.Content
	updatedMissing.LocationMissing = locationsmissing.LocationMissing{
		Name: missingDomain.LocationMissing.Name,
		Lat:  missingDomain.LocationMissing.Lat,
		Lng:  missingDomain.LocationMissing.Lng,
	}
	mr.conn.Where("id=?", updatedMissing.LocationMissing.ID).UpdateColumns(missing.LocationMissing)
	for _, hashtag := range missingDomain.Hashtags {
		updatedMissing.Hashtags = append(updatedMissing.Hashtags, hashtags.Hashtag{
			Name: hashtag.Name,
		})
	}
	updatedMissing.IsFound = missingDomain.IsFound
	mr.conn.Save(&updatedMissing)
	return updatedMissing.ToDomain()
}

func (mr *missingRepository) Delete(id string) bool {
	var missing missings.Domain = mr.GetByID(id)

	deletedMissing := FromDomain(&missing)
	if result := mr.conn.Unscoped().Delete(&deletedMissing); result.RowsAffected == 0 {
		return false
	}
	return true
}

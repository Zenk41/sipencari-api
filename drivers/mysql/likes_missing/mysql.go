package likesmissing

import (
	likesmissing "sipencari-api/businesses/likes_missing"

	"gorm.io/gorm"
)

type likesMissingRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) likesmissing.Repository {
	return &likesMissingRepository{
		conn: conn,
	}
}

func (lmr *likesMissingRepository) GetAll() []likesmissing.Domain {
	var rec []LikeMissing
	lmr.conn.Preload("User").Find(&rec)
	likeMissingDomain := []likesmissing.Domain{}
	for _, likeMissing := range rec {
		likeMissingDomain = append(likeMissingDomain, likeMissing.ToDomain())
	}
	return likeMissingDomain
}

func (lmr *likesMissingRepository) GetByID(idLike string, idMissing string) likesmissing.Domain {
	likeMissing := LikeMissing{}
	lmr.conn.
		Where("user_id=? AND missing_id=?", idLike, idMissing).
		Preload("User").
		First(&likeMissing)
	return likeMissing.ToDomain()
}

func (lmr *likesMissingRepository) Like(idUser string, idMissing string, likeDomain *likesmissing.Domain) likesmissing.Domain {
	var likesmissing likesmissing.Domain = lmr.GetByID(idUser, idMissing)
	rec := FromDomain(likeDomain)
	rec.MissingID = idMissing
	if likesmissing.UserID != idUser {
		lmr.conn.Preload("User").Create(&rec)
		return rec.ToDomain()
	}
	return rec.ToDomain()
}

func (lmr *likesMissingRepository) Unlike(idUser string, idMissing string) bool {
	var likesmissing likesmissing.Domain = lmr.GetByID(idUser, idMissing)
	UnLike := FromDomain(&likesmissing)
	if result := lmr.conn.
		Unscoped().
		Where("user_id=?", idUser).
		Delete(&UnLike); result.RowsAffected == 0 {
		return false
	}
	return true
}

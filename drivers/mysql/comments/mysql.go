package comments

import (
	"sipencari-api/businesses/comments"
	locationscomment "sipencari-api/drivers/mysql/locations_comment"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type commentRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) comments.Repository {
	return &commentRepository{
		conn: conn,
	}
}

func (comr *commentRepository) GetAll(idMissing string) []comments.Domain {
	var rec []Comment
	comr.conn.
		Preload("User").
		Preload("LocationComment").
		Preload("LikeComments.User").
		Find(&rec, "missing_id=?", idMissing)
	commentDomain := []comments.Domain{}
	for _, comment := range rec {
		commentDomain = append(commentDomain, comment.ToDomain())
	}
	return commentDomain
}

func (comr *commentRepository) GetByID(idMissing string, id string) comments.Domain {
	comment := Comment{}
	comr.conn.
		Preload("User").
		Preload("LocationComment").
		Preload("LikeComments.User").
		Where("missing_id=? AND id=?", idMissing, id).
		First(&comment)
	return comment.ToDomain()
}

func (comr *commentRepository) Create(idUser string, idMissing string, commentDomain *comments.Domain) comments.Domain {
	rec := FromDomain(commentDomain)
	rec.UserID = idUser
	rec.MissingID = idMissing
	rec.LocationComment.ID = uuid.NewString()
	result := comr.conn.
		Preload("User").
		Preload("LocationComment").
		Preload("LikeComments.User").
		Create(&rec)
	result.Last(&rec)
	return rec.ToDomain()
}

func (comr *commentRepository) Update(idUser string, idMissing string, idComment string, commentDomain *comments.Domain) comments.Domain {
	var comment comments.Domain = comr.GetByID(idMissing, idComment)
	updatedComment := FromDomain(&comment)
	updatedComment.Message = commentDomain.Message
	updatedComment.LocationComment = locationscomment.LocationComment{
		Name: commentDomain.LocationComment.Name,
		Lat:  commentDomain.LocationComment.Lat,
		Lng:  commentDomain.LocationComment.Lng,
	}
	comr.conn.Save(&updatedComment)
	return updatedComment.ToDomain()
}

func (comr *commentRepository) Delete(idMissing string, id string) bool {
	var comment comments.Domain = comr.GetByID(idMissing, id)

	deletedComment := FromDomain(&comment)
	if result := comr.conn.Unscoped().Delete(&deletedComment); result.RowsAffected == 0 {
		return false
	}
	return true
}

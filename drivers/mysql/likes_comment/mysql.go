package likescomment

import (
	likescomment "sipencari-api/businesses/likes_comment"

	"gorm.io/gorm"
)

type likesCommentRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) likescomment.Repository {
	return &likesCommentRepository{
		conn: conn,
	}
}

func (lcr *likesCommentRepository) GetAll() []likescomment.Domain {
	var rec []LikeComment
	lcr.conn.Preload("User").Find(&rec)
	likeCommentDomain := []likescomment.Domain{}
	for _, likeComment := range rec {
		likeCommentDomain = append(likeCommentDomain, likeComment.ToDomain())
	}
	return likeCommentDomain
}

func (lcr *likesCommentRepository) GetByID(idLike string, idMissing string, idComment int) likescomment.Domain {
	likeComment := LikeComment{}
	lcr.conn.
		Where("user_id=? AND comment_id=?", idLike, idComment).
		Preload("User").
		First(&likeComment)
	return likeComment.ToDomain()
}

func (lcr *likesCommentRepository) Like(idUser string, idMissing string,  idComment int,likeDomain *likescomment.Domain) likescomment.Domain {
	var likescomment likescomment.Domain = lcr.GetByID(idUser, idMissing, idComment)
	rec := FromDomain(likeDomain)
	rec.CommentID = uint(idComment)
	if likescomment.UserID != idUser {
		lcr.conn.Preload("User").
		Create(&rec)
		return rec.ToDomain()
	}
	return rec.ToDomain()
}

func (lcr *likesCommentRepository) Unlike(idUser string, idMissing string,idComment int) bool {
	var likescomment likescomment.Domain = lcr.GetByID(idUser, idMissing, idComment)
	UnLike := FromDomain(&likescomment)
	if result := lcr.conn.
		Unscoped().
		Where("user_id=?", idUser).
		Delete(&UnLike); result.RowsAffected == 0 {
		return false
	}
	return true
}

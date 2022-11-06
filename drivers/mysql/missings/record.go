package missings

import (
	busComment "sipencari-api/businesses/comments"
	busHashtag "sipencari-api/businesses/hashtags"
	busLike "sipencari-api/businesses/likes_missing"
	"sipencari-api/businesses/missings"
	mysqlComment "sipencari-api/drivers/mysql/comments"
	mysqlHashtag "sipencari-api/drivers/mysql/hashtags"

	// likescomment "sipencari-api/drivers/mysql/likes_comment"
	mysqlLike "sipencari-api/drivers/mysql/likes_missing"
	locationscomment "sipencari-api/drivers/mysql/locations_comment"
	locations "sipencari-api/drivers/mysql/locations_missing"

	"sipencari-api/drivers/mysql/categories"
	"sipencari-api/drivers/mysql/users"
	"time"

	"gorm.io/gorm"
)

type Missing struct {
	ID              string                    `json:"id" gorm:"size:100;primaryKey"`
	Title           string                    `json:"title"`
	Content         string                    `json:"content"`
	Image           string                    `json:"image"`
	User            users.User                `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID          string                    `json:"user_id" gorm:"size:100"`
	LocationMissing locations.LocationMissing `json:"location_missing" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CategoryID      uint                      `json:"category_id"`
	Category        categories.Category       `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Hashtags        []mysqlHashtag.Hashtag    `json:"hashtags" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LikeMissings    []mysqlLike.LikeMissing   `json:"like_missings" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments        []mysqlComment.Comment    `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsFound         bool                      `json:"is_found"`
	CreatedAt       time.Time                 `json:"created_at"`
	UpdatedAt       time.Time                 `json:"updated_at"`
	DeletedAt       gorm.DeletedAt            `json:"deleted_at"`
}

func FromDomain(domain *missings.Domain) *Missing {
	var hashtagsData []mysqlHashtag.Hashtag
	hashtagFromDomain := domain.Hashtags
	for _, hashtag := range hashtagFromDomain {
		hashtagsData = append(hashtagsData, mysqlHashtag.Hashtag{
			Name: hashtag.Name,
		})
	}

	var likesMissingData []mysqlLike.LikeMissing
	likesmissingFromDomain := domain.LikeMissings
	for _, likemissing := range likesmissingFromDomain {
		likesMissingData = append(likesMissingData, mysqlLike.LikeMissing{
			UserID:    likemissing.UserID,
			MissingID: likemissing.MissingID,
		})
	}

	var commentsData []mysqlComment.Comment
	commentsFromDomain := domain.Comments
	for _, comment := range commentsFromDomain {
		commentsData = append(commentsData, mysqlComment.Comment{
			ID:        comment.ID,
			UserID:    comment.UserID,
			Image:     comment.Image,
			MissingID: comment.MissingID,
			LocationComment: locationscomment.LocationComment{
				ID:        comment.LocationComment.ID,
				Name:      comment.LocationComment.Name,
				Lat:       comment.LocationComment.Lat,
				Lng:       comment.LocationComment.Lng,
				CommentID: comment.LocationComment.CommentID,
			},
		})
	}

	return &Missing{
		ID:      domain.ID,
		Title:   domain.Title,
		Content: domain.Content,
		Image:   domain.Image,
		UserID:  domain.UserID,
		LocationMissing: locations.LocationMissing{
			ID:        domain.LocationMissing.ID,
			Name:      domain.LocationMissing.Name,
			Lat:       domain.LocationMissing.Lat,
			Lng:       domain.LocationMissing.Lng,
			MissingID: domain.LocationMissing.MissingID,
		},
		CategoryID:   domain.CategoryID,
		Hashtags:     hashtagsData,
		LikeMissings: likesMissingData,
		Comments:     commentsData,
		IsFound:      domain.IsFound,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
	}
}

func (rec *Missing) ToDomain() missings.Domain {
	var commentsFromDomain []busComment.Domain
	for _, comment := range rec.Comments {
		commentsFromDomain = append(commentsFromDomain, comment.ToDomain())
	}
	var hashtagsFromDomain []busHashtag.Domain
	for _, hashtag := range rec.Hashtags {
		hashtagsFromDomain = append(hashtagsFromDomain, hashtag.ToDomain())
	}

	var likesFromDomain []busLike.Domain
	for _, like := range rec.LikeMissings {
		likesFromDomain = append(likesFromDomain, like.ToDomain())
	}

	return missings.Domain{
		ID:              rec.ID,
		Title:           rec.Title,
		Content:         rec.Content,
		Image:           rec.Image,
		UserID:          rec.UserID,
		UserName:        rec.User.Name,
		LocationMissing: rec.LocationMissing.ToDomain(),
		CategoryID:      rec.CategoryID,
		CategoryName:    rec.Category.Name,
		Hashtags:        hashtagsFromDomain,
		LikeMissings:    likesFromDomain,
		Comments:        commentsFromDomain,
		IsFound:         rec.IsFound,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
		DeletedAt:       rec.DeletedAt,
	}
}

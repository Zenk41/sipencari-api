package response

import (
	"sipencari-api/businesses/hashtags"
	likesmissing "sipencari-api/businesses/likes_missing"
	"sipencari-api/businesses/missings"
	resComment "sipencari-api/controllers/comments/response"
	resHashtag "sipencari-api/controllers/hashtags/response"
	resLikeComment "sipencari-api/controllers/likes_comment/response"
	resLikeMissing "sipencari-api/controllers/likes_missing/response"
	"sipencari-api/controllers/locations_comment/response"
	resLocation "sipencari-api/controllers/locations_missing/response"
	"time"

	"gorm.io/gorm"
)

type Missing struct {
	ID              string                       `json:"id" gorm:"primaryKey"`
	Title           string                       `json:"title"`
	Content         string                       `json:"content"`
	Image           string                       `json:"image"`
	UserName        string                       `json:"user_name"`
	UserID          string                       `json:"user_id"`
	LocationMissing resLocation.LocationMissing  `json:"location_missing"`
	CategoryName    string                       `json:"category_name"`
	CategoryID      uint                         `json:"category_id"`
	Hashtags        []resHashtag.Hashtag         `json:"hashtags"`
	LikeMissings    []resLikeMissing.LikeMissing `json:"like_missings"`
	Comments        []resComment.Comment         `json:"comments"`
	IsFound         bool                         `json:"is_found"`
	CreatedAt       time.Time                    `json:"created_at"`
	UpdatedAt       time.Time                    `json:"updated_at"`
	DeletedAt       gorm.DeletedAt               `json:"deleted_at"`
}

func FromDomain(domain missings.Domain) Missing {
	var hashtagsData []resHashtag.Hashtag
	hashtagToDomain := domain.Hashtags
	for _, hashtag := range hashtagToDomain {
		hashtagsData = append(hashtagsData, resHashtag.FromDomain(hashtags.Domain{
			Name: hashtag.Name,
		}))
	}

	var likesMissingData []resLikeMissing.LikeMissing
	likesmissingFromDomain := domain.LikeMissings
	for _, likemissing := range likesmissingFromDomain {
		likesMissingData = append(likesMissingData, resLikeMissing.FromDomain(likesmissing.Domain{
			UserID: likemissing.UserID,
			UserName: likemissing.UserName,
			MissingID: likemissing.MissingID,
		}))
	}

	var commentsData []resComment.Comment
	var likesCommentData []resLikeComment.LikeComment
	commentsFromDomain := domain.Comments
	var likesCommentFromDomain []resLikeComment.LikeComment
	for _, comment := range commentsFromDomain {
		for _, likecomment := range likesCommentFromDomain {
			likesCommentData = append(likesCommentData, resLikeComment.LikeComment{
				UserID:   likecomment.UserID,
				UserName: likecomment.UserName,
			})
		}
		commentsData = append(commentsData, resComment.Comment{
			ID:       comment.ID,
			UserID:   comment.UserID,
			UserName: comment.UserName,
			Message:  comment.Message,
			Image:    comment.Image,
			LocationComment: response.LocationComment{
				ID:   comment.LocationComment.ID,
				Name: comment.LocationComment.Name,
				Lat:  comment.LocationComment.Lat,
				Lng:  comment.LocationComment.Lng,
			},
			LikeComments: likesCommentData,
		})
	}
	return Missing{
		ID:       domain.ID,
		Title:    domain.Title,
		Content:  domain.Content,
		Image:    domain.Image,
		UserName: domain.UserName,
		UserID:   domain.UserID,
		LocationMissing: resLocation.LocationMissing{
			ID:   domain.LocationMissing.ID,
			Name: domain.LocationMissing.Name,
			Lat:  domain.LocationMissing.Lat,
			Lng:  domain.LocationMissing.Lng,
		},
		CategoryName: domain.CategoryName,
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

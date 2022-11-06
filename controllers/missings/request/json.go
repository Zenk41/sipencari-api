package request

import (
	"sipencari-api/businesses/comments"
	"sipencari-api/businesses/hashtags"
	likesmissing "sipencari-api/businesses/likes_missing"
	locationsmissing "sipencari-api/businesses/locations_missing"
	"sipencari-api/businesses/missings"
	reqComment "sipencari-api/controllers/comments/request"
	reqHashtag "sipencari-api/controllers/hashtags/request"
	reqLike "sipencari-api/controllers/likes_missing/request"
	reqLocation "sipencari-api/controllers/locations_missing/request"

	"github.com/go-playground/validator"
)

type Missing struct {
	Title           string                      `json:"title" form:"title" validate:"required"`
	Content         string                      `json:"content" form:"content" validate:"required"`
	Image           string                      `json:"image" form:"image"`
	LocationMissing reqLocation.LocationMissing `json:"location_missing" form:"location_missing" validate:"required"`
	CategoryID      uint                        `json:"category_id" form:"category_id" validate:"required"`
	Hashtags        []reqHashtag.Hashtag        `json:"hashtags" form:"hashtags"`
	LikesMissings   []reqLike.LikeMissing       `json:"like_missings" form:"like_missings"`
	Comments        []reqComment.Comment        `json:"comments" form:"comments"`
	IsFound         bool                        `json:"is_found" form:"is_found" `
}

func (req *Missing) ToDomain() *missings.Domain {
	var hashtags []hashtags.Domain
	for _, hashtag := range req.Hashtags {
		hashtags = append(hashtags, *hashtag.ToDomain())
	}

	var likesmissing []likesmissing.Domain
	for _, likemissing := range req.LikesMissings {
		likesmissing = append(likesmissing, *likemissing.ToDomain())
	}

	var comments []comments.Domain
	for _, comment := range req.Comments {
		comments = append(comments, *comment.ToDomain())
	}

	return &missings.Domain{
		Title:   req.Title,
		Content: req.Content,
		Image:   req.Image,
		LocationMissing: locationsmissing.Domain{
			Name: req.LocationMissing.Name,
			Lat:  req.LocationMissing.Lat,
			Lng:  req.LocationMissing.Lng,
		},
		CategoryID:   req.CategoryID,
		Hashtags:     hashtags,
		LikeMissings: likesmissing,
		Comments:     comments,
		IsFound:      req.IsFound,
	}
}

func (req *Missing) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

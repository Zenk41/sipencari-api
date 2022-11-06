package request

import "sipencari-api/businesses/hashtags"

type Hashtag struct {
	Name string `json:"name" form:"name"`
}

func (req *Hashtag) ToDomain() *hashtags.Domain {
	return &hashtags.Domain{
		Name: req.Name,
	}
}
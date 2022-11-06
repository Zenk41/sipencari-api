package request

import locations "sipencari-api/businesses/locations_comment"

type LocationComment struct {
	Name string  `json:"name" form:"name"`
	Lat  string `json:"lat" form:"lat"`
	Lng  string `json:"lng" form:"lng"`
}

func (req *LocationComment) ToDomain() *locations.Domain {
	return &locations.Domain{
		Name: req.Name,
		Lat:  req.Lat,
		Lng:  req.Lng,
	}
}

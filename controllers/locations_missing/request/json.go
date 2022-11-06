package request

import locations "sipencari-api/businesses/locations_missing"

type LocationMissing struct {
	Name string  `json:"name" form:"name"`
	Lat  string `json:"lat" form:"lat"`
	Lng  string `json:"lng" form:"lng"`
}

func (req *LocationMissing) ToDomain() *locations.Domain {
	return &locations.Domain{
		Name: req.Name,
		Lat:  req.Lat,
		Lng:  req.Lng,
	}
}

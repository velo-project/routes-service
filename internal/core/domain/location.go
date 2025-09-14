package domain

type Location struct {
	ID  *int    `json:"id,omitempty"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func NewLocation(id *int, lat, lng float64) *Location {
	return &Location{
		ID:  id,
		Lat: lat,
		Lng: lng,
	}
}

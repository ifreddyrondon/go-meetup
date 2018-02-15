package point

import (
	"encoding/json"
)

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// START OMIT
type pointJSON struct { // HL
	Lat       float64 `json:"lat"`
	Latitude  float64 `json:"latitude"`
	Lng       float64 `json:"lng"`
	Longitude float64 `json:"longitude"`
}

func (po *Point) UnmarshalJSON(data []byte) error { // HL
	var model pointJSON
	if err := json.Unmarshal(data, &model); err != nil {
		panic(err)
	}

	// getLat and getLng are helper func to get filled values
	lat, lng := getLat(&model), getLng(&model)
	point := &Point{Lat: lat, Lng: lng}
	*po = *point
	return nil
}

// END OMIT

func getLat(model *pointJSON) float64 {
	if model.Lat == 0 {
		return model.Latitude
	}
	return model.Lat
}

func getLng(model *pointJSON) float64 {
	if model.Lng == 0 {
		return model.Longitude
	}
	return model.Lng
}

package geometry

import "math"

type Coords struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

/**
Haversine formula implementation to calculate distance between two coordinates on a sphere
*/
func (c Coords) Dist(c2 Coords) float64 {
	delta := getDelta(c, c2)
	dLat := toRad(delta.Lat)
	dLon := toRad(delta.Lon)
	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(toRad(c.Lat))*math.Cos(toRad(c2.Lat))*math.Pow(math.Sin(dLon/2), 2)
	return 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a)) * 6372.8 // Earth's radius in km
}

/**
Calculate bounding box for the database index with approximates and safety padding
*/
func (c Coords) GetBoundingBox(r float64) (bottomLeft, topRight Coords) {
	lat := r / 111.699 * 2
	lon := r / 111.321 * 2
	return Coords{Lat: c.Lat - lat, Lon: c.Lon - lon}, Coords{Lat: c.Lat + lat, Lon: c.Lon + lon}
}

func getDelta(c1, c2 Coords) Coords {
	return Coords{
		Lat: c2.Lat - c1.Lat,
		Lon: c2.Lon - c1.Lon,
	}
}

func toRad(f float64) float64 {
	return f * math.Pi / 180
}

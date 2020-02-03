package database

import (
	"airport/cmd/geometry"
	"encoding/json"
)

type Airport struct {
	Coords   geometry.Coords `json:"coords"`
	Name     string          `json:"name"`
	Distance float64         `json:"distance"`
}

type Airports []Airport

func (a Airports) Len() int           { return len(a) }
func (a Airports) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Airports) Less(i, j int) bool { return a[i].Distance < a[j].Distance }
func (a Airports) ToJson() string {
	str := "["
	for i, airport := range a {
		bytes, _ := json.Marshal(airport)
		str = str + string(bytes)
		if i+1 < len(a) {
			str = str + ","
		}
	}
	str = str + "]"
	return str
}

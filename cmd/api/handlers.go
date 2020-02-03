package api

import (
	"airport/cmd/database"
	"airport/cmd/geometry"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

/**
Returns JSON list of airports matching the criteria
Happy path implementation, error handling is omitted for simplicity
*/
func ListAirports(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	lat, _ := strconv.ParseFloat(q["lat"][0], 64)
	lon, _ := strconv.ParseFloat(q["lon"][0], 64)
	radius, _ := strconv.ParseFloat(q["r"][0], 64)
	log.Printf("GET /list - lat=%f lon=%f r=%f", lat, lon, radius)

	airports := database.GetAirports(geometry.Coords{Lat: lat, Lon: lon}, radius)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = fmt.Fprintf(w, airports.ToJson())
}

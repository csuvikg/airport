package database

import (
	"airport/cmd/geometry"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
)

const geoSearchUrl = "https://mikerhodes.cloudant.com/airportdb/_design/view1/_search/geo"
const limit = 200

func GetAirports(center geometry.Coords, r float64) Airports {
	data := loadData(center, r)
	filteredData := filterData(data, center, r)
	sort.Sort(filteredData)
	return filteredData
}

/**
Filters coordinates by distance
*/
func filterData(data []dataRow, center geometry.Coords, r float64) Airports {
	valid := make([]Airport, 0, len(data))
	for _, row := range data {
		coord := geometry.Coords{
			Lat: row.Fields.Lat,
			Lon: row.Fields.Lon,
		}
		dist := coord.Dist(center)
		if dist <= r {
			valid = append(valid, Airport{
				Coords:   coord,
				Name:     row.Fields.Name,
				Distance: dist,
			})
		}
	}
	return valid
}

func getQueryString(center geometry.Coords, r float64) string {
	c1, c2 := center.GetBoundingBox(r)
	query := fmt.Sprintf("lon:[%f TO %f] AND lat:[%f TO %f]", c1.Lon, c2.Lon, c1.Lat, c2.Lat)
	return url.QueryEscape(query)
}

/**
Returns aggregated data from Cloudant API
Cloudant index API returns pages by bookmark
Assumes correctness of service and data, no error handling implemented
*/
func loadData(center geometry.Coords, r float64) []dataRow {
	preparedUrl := fmt.Sprintf("%s?limit=%d&q=%s", geoSearchUrl, limit, getQueryString(center, r))

	var results []dataRow
	requestUrl := preparedUrl
	for {
		response, _ := http.Get(requestUrl)
		contents, _ := ioutil.ReadAll(response.Body)
		var airportsResponse apiResponse
		_ = json.Unmarshal(contents, &airportsResponse)
		_ = response.Body.Close()

		if len(airportsResponse.Rows) > 0 {
			results = append(results, airportsResponse.Rows...)
			requestUrl = fmt.Sprintf("%s&bookmark=%s", preparedUrl, airportsResponse.Bookmark)
		} else {
			break
		}
	}
	return results
}

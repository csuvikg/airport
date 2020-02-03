package database

type dataField struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Name string  `json:"name"`
}

type dataRow struct {
	Id     string    `json:"id"`
	Order  []float64 `json:"order"`
	Fields dataField `json:"fields"`
}

type apiResponse struct {
	RowsCount int       `json:"total_rows"`
	Bookmark  string    `json:"bookmark"`
	Rows      []dataRow `json:"rows"`
}

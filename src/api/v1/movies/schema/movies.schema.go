package movies_schema

type Movie struct {
	Title    string `json:"title"`
	Year     string `json:"release_year"`
	Director string `json:"director"`
	Genre    string `json:"genre"`
}

const SCHEMA_NAME = "movies"

// [
// 	{
// 		"title": "The Green Hornet",
// 		"release_year": "1993"
// 	}
// ]

package movies_schema

type Movie struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string `json:"title"`
	Year     string `json:"year"`
	Director string `json:"director"`
	Genre    string `json:"genre"`
}

const SCHEMA_NAME = "movies"

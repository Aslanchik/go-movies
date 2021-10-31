package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title"`
	Year     string             `json:"year"`
	Director string             `json:"director"`
	Genre    string             `json:"genre"`
}

type DeleteRequest struct {
	ID string `json:"id" binding:"required,gte=1"`
}

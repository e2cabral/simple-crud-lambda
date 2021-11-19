package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	PublicationDate string             `bson:"publication_date" json:"publication_date"`
	Author          string             `bson:"author" json:"author"`
	Pages           int16              `bson:"pages" json:"pages"`
	Subject         string             `bson:"subject" json:"subject"`
	Characters      []string           `bson:"characters" json:"characters"`
	Genre           string             `bson:"genre" json:"genre"`
	Price           float32            `bson:"price" json:"price"`
	Rating          int8               `bson:"rating" json:"rating"`
}

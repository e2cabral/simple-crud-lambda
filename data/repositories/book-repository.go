package repositories

import (
	"aws-lambda-simple-example/domain"
	"aws-lambda-simple-example/infra/database"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type BookRepository struct {
	*mongo.Collection
}

func NewBookRepository() *BookRepository {
	collection := database.Connect().Database("book_store_app").Collection("books")
	return &BookRepository{collection}
}

func (b *BookRepository) Create(book domain.Book) []byte {
	formatted := domain.Book{
		ID:              primitive.NewObjectID(),
		PublicationDate: book.PublicationDate,
		Author:          book.Author,
		Pages:           book.Pages,
		Subject:         book.Subject,
		Characters:      book.Characters,
		Genre:           book.Genre,
		Price:           book.Price,
		Rating:          book.Rating,
	}
	result, err := b.InsertOne(context.TODO(), formatted)
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func (b *BookRepository) Get() []byte {
	var results []*domain.Book
	filterOptions := options.Find()
	filterOptions.SetLimit(100)

	result, err := b.Find(context.TODO(), bson.D{}, filterOptions)
	if err != nil {
		log.Fatal(err)
	}

	for result.Next(context.TODO()) {
		var book domain.Book

		err := result.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &book)
	}

	data, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func (b *BookRepository) GetById(id string) []byte {
	var book *domain.Book

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	result := b.FindOne(
		context.TODO(),
		bson.D{{"_id", objectID}},
	)

	if err := result.Decode(&book); err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func (b *BookRepository) Update(id string, book domain.Book) []byte {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", objectId}}
	itemUpdated := bson.D{{"$set", &book}}

	result, err := b.UpdateOne(context.TODO(), filter, itemUpdated)
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func (b *BookRepository) Delete(id string) []byte {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	result, err := b.DeleteOne(context.TODO(), bson.D{{"_id", objectId}})
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

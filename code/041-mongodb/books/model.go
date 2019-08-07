package books

import (
	"context"
	"errors"
	"log"

	"github.com/dnilosek/learning-web-dev/code/041-mongodb/config"
	"go.mongodb.org/mongo-driver/bson"
)

type Book struct {
	ISBN   string  `bson:isbn`
	Title  string  `bson:title`
	Author string  `bson:author`
	Price  float64 `bson:price`
}

var ErrRequest = errors.New("Reuest not valid")
var ErrAccept = errors.New("Reuest not acceptable")
var ErrNotFound = errors.New("Not found")
var ErrInternal = errors.New("Internal Error")

func GetOne(isbn string) (Book, error) {
	bk := Book{}
	if isbn == "" {
		return bk, ErrRequest
	}
	err := config.DB.FindOne(context.TODO(), bson.D{{"isbn", isbn}}).Decode(&bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func DeleteOne(isbn string) error {
	if isbn == "" {
		return ErrRequest
	}
	_, err := config.DB.DeleteOne(context.TODO(), bson.D{{"isbn", isbn}})
	if err != nil {
		return err
	}
	return nil
}

func InsertOne(bk Book) error {
	// insert
	_, err := config.DB.InsertOne(context.TODO(), bk)
	if err != nil {
		return err
	}
	return nil
}

func UpdateOne(bk Book) error {
	// insert
	bDoc, err := toDoc(bk)
	if err != nil {
		log.Println("Cannot toDoc")
		return err
	}
	update := bson.D{
		{"$set", bDoc},
	}
	_, err = config.DB.UpdateOne(context.TODO(), bson.D{{"isbn", bk.ISBN}}, update)
	if err != nil {
		log.Println("Cannot update")
		return err
	}
	return nil
}

func GetAll() ([]Book, error) {
	bks := make([]Book, 0)
	cur, err := config.DB.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return bks, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var bk Book
		err := cur.Decode(&bk)
		if err != nil {
			return bks, err
		}
		bks = append(bks, bk)
	}
	return bks, nil
}

func toDoc(v Book) (bson.D, error) {
	var doc bson.D
	data, err := bson.Marshal(v)
	if err != nil {
		log.Println("Cannot marshal")
		return doc, err
	}
	err = bson.Unmarshal(data, &doc)
	return doc, err
}

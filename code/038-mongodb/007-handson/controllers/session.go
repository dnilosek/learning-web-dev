package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/dnilosek/learning-web-dev/code/038-mongodb/007-handson/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionController struct {
	collection *mongo.Collection
}

func NewSessionController(c *mongo.Client) *SessionController {
	return &SessionController{c.Database("go-web-dev-db").Collection("sessions")}
}

func (sc *SessionController) GetSession(sessionId string) (models.Session, error) {
	var s models.Session
	// Create hex from id and find it
	idHex, err := primitive.ObjectIDFromHex(sessionId)
	if err != nil {
		return s, err
	}
	err = sc.collection.FindOne(context.TODO(), bson.D{{"_id", idHex}}).Decode(&s)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (sc *SessionController) DeleteSession(sessionId string) error {
	// Create hex from id and find it
	idHex, err := primitive.ObjectIDFromHex(sessionId)
	if err != nil {
		return err
	}
	_, err = sc.collection.DeleteOne(context.TODO(), bson.D{{"_id", idHex}})
	if err != nil {
		return err
	}
	return nil
}

// Create session for a given user id, return session id
func (sc *SessionController) CreateSession(s models.Session) (string, error) {
	var sId string
	insertResult, err := sc.collection.InsertOne(context.TODO(), s)
	if err != nil {
		return sId, err
	}

	fmt.Sprint(sId, insertResult.InsertedID.(primitive.ObjectID).Hex)
	sId = fmt.Sprint(insertResult.InsertedID.(primitive.ObjectID).Hex())
	return sId, nil
}

// Update session activity in model and db
func (sc *SessionController) UpdateActivity(sessionId string, s models.Session) (models.Session, error) {
	s.LastActivity = time.Now()
	idHex, err := primitive.ObjectIDFromHex(sessionId)
	if err != nil {
		return s, err
	}
	// Update
	updateVal := bson.D{
		{"$set",
			bson.D{{"lastactivity", s.LastActivity}}},
	}
	_, err = sc.collection.UpdateOne(context.TODO(), bson.D{{"_id", idHex}}, updateVal)
	if err != nil {
		return s, err
	}
	return s, nil
}

// Clean sessions
func (sc *SessionController) CleanSessions() (int64, error) {

	filter := bson.M{"lastactivity": bson.M{
		"$lt": time.Now().Add(-30 * time.Second),
	}}

	// update documents
	res, err := sc.collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	// Clear results
	return res.DeletedCount, nil
}

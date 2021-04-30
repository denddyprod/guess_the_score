package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func NewPredictionModel(services *Services) PredictionModel {
	return &predictionMongo{services}
}

type Prediction struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	MatchId     primitive.ObjectID             `json:"match_id" bson:"match_id,omitempty"`
	UserId primitive.ObjectID             `json:"user_id" bson:"user_id,omitempty"`
	TeamAScore     uint32             `json:"teamA_score" bson:"teamA_score,omitempty"`
	TeamBScore      uint32             `json:"teamB_score" bson:"teamB_score,omitempty"`
}

type PredictionModel interface {
	FindByUserId(userId primitive.ObjectID) ([]Prediction, error)
}

var _ PredictionModel = &predictionMongo{}

type predictionMongo struct {
	servs *Services
}

func (pm predictionMongo) FindByUserId(userId primitive.ObjectID) ([]Prediction, error) {
	var results []Prediction

	predictionCollection := pm.servs.db.Collection("predictions")
	filter := bson.D{{"user_id", userId}}

	cursor, err := predictionCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	fmt.Printf("Extracted all predictions with filter: %+v\n", results)
	return results, nil
}

func (pm predictionMongo) Create(prediction *Prediction) error {
	predictionCollection := pm.servs.db.Collection("predictions")

	res, err := predictionCollection.InsertOne(context.TODO(), prediction)
	if err != nil {
		return err
	}

	log.Println("Created match with id=", res.InsertedID)
	return nil
}

func (pm predictionMongo) Update(prediction *Prediction) error {
	predictionCollection := pm.servs.db.Collection("predictions")

	filter := bson.D{{"_id", prediction.Id}}

	updateResult, err := predictionCollection.ReplaceOne(context.TODO(), filter, prediction)
	if err != nil {
		return err
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return nil
}

func (pm predictionMongo) Delete(id primitive.ObjectID) error {
	predictionCollection := pm.servs.db.Collection("predictions")

	filter := bson.M{"_id": id}

	_, err := predictionCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	fmt.Printf("Deleted object with _id: %v \n", id)
	return nil
}
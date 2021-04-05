package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

// User represents the user model stored in our database
type User struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Username        string             `json:"username" bson:"username"`
	PasswordHash    string             `json:"password_hash" bson:"password_hash"`
	Password        string             `json:"-" bson:"-"`
	JWToken         string             `json:"jwt_token" bson:"jwt_token"`
	ActivationToken string             `json:"activation_token" bson:"activation_token"`
	AccessRights    AccessRights       `json:"access_rights" bson:"access_rights"`
	IsActive        bool               `json:"is_active" bson:"is_active"`
	Score           int                `json:"score" bson:"score"`
}

type AccessRights struct {
	Read  bool `json:"read" bson:"read"`
	Write bool `json:"write" bson:"write"`
	Edit  bool `json:"edit" bson:"edit"`
}

type UserModel interface {
	// Methods for querying for single user
	FindById(id primitive.ObjectID) (*User, error)
	FindByUsername(username string) (*User, error)

	// This function validate all fields for create and update
	Validate(user *User, action string) error

	// Methods for altering user
	Create(user *User) error
	Update(user *User) error
	Delete(id primitive.ObjectID) error
}

var _ UserModel = &userValidator{}
var _ UserModel = &userMongo{}

type userMongo struct {
	servs *Services
}

func (ug *userMongo) FindById(id primitive.ObjectID) (*User, error) {
	usersCollection := ug.servs.db.Collection("users")
	var resultUser User

	filter := bson.D{{"_id", id}}

	err := usersCollection.FindOne(context.TODO(), filter).Decode(&resultUser)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Found a single document: %+v\n", resultUser)

	return &resultUser, nil
}

func (ug *userMongo) FindByUsername(username string) (*User, error) {
	usersCollection := ug.servs.db.Collection("users")
	var resultUser User

	filter := bson.D{{"username", username}}

	err := usersCollection.FindOne(context.TODO(), filter).Decode(&resultUser)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Found a single document: %+v\n", resultUser)

	return &resultUser, nil
}

func (ug *userMongo) Validate(user *User, action string) error {

	return nil
}

func (ug *userMongo) Create(user *User) error {
	usersCollection := ug.servs.db.Collection("users")

	res, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	log.Println("Created user with id=", res.InsertedID)

	return nil
}

func (ug *userMongo) Update(user *User) error {
	usersCollection := ug.servs.db.Collection("users")

	filter := bson.D{{"_id", user.Id}}

	updateResult, err := usersCollection.ReplaceOne(context.TODO(), filter, user)
	if err != nil {
		return err
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return nil
}

func (ug *userMongo) Delete(id primitive.ObjectID) error {
	usersCollection := ug.servs.db.Collection("users")

	filter := bson.M{"_id": id}

	_, err := usersCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	fmt.Printf("Deleted object with _id: %v \n", id)
	return nil
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entry struct {
	ID          primitive.ObjectID `bson:"id"`
	Dish        *string            `json:"dish"`
	Fats        *float64           `json:"fats"`
	Ingredients *string            `json:"ingredients"`
	Calories    *string            `json:"calories"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Desc string             `bson:"desc"`
	Kg   int32              `bson:"Kg"`
}

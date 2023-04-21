package service

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Yfleet/shared_proto/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DelItem(client *mongo.Client, rq *pb.DelItemRequest) bool {
	collection := client.Database("Warehouse").Collection("Items")

	objectID, err := primitive.ObjectIDFromHex(rq.ID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	filter := bson.M{"_id": objectID}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Result:", result)
	return true
}

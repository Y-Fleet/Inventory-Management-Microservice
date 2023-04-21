package service

import (
	"context"
	"log"
	"time"

	pb "github.com/Y-Fleet/Grpc-Api/api"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddItemToDb(rq *pb.AddItemRequest, client *mongo.Client) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.Database("Warehouse").Collection("Items")
	items := protoToStruct(rq)
	_, err := collection.InsertOne(ctx, items)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

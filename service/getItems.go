package service

import (
	"context"
	models "inventory/models"
	"log"
	"time"

	proto "github.com/Yfleet/shared_proto/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetItem(client *mongo.Client) *proto.GetItemResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.Database("Warehouse").Collection("Items")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var items []models.Item
	if err = cursor.All(ctx, &items); err != nil {
		log.Fatal(err)
	}
	return ItemsToProto(items, err)
}

func GetInventory(client *mongo.Client, rq *proto.GetInventoryRequest) ([]models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var itemIDs []primitive.ObjectID
	for _, itemID := range rq.ID {
		objectID, err := primitive.ObjectIDFromHex(itemID.ID)
		if err != nil {
			log.Printf("Error converting item ID to ObjectID: %v\n", err)
			continue
		}
		itemIDs = append(itemIDs, objectID)
	}
	if len(itemIDs) == 0 {
		return []models.Item{}, nil
	}

	filter := bson.M{"_id": bson.M{"$in": itemIDs}}
	cursor, err := client.Database("Warehouse").Collection("Items").Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []models.Item
	if err := cursor.All(ctx, &items); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return items, nil
}

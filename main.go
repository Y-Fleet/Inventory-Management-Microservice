package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	service "inventory/service"

	pb "github.com/Y-Fleet/Grpc-Api/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedInventoryServiceServer
}

func (s *server) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	client, cancel := ConToDb()
	if service.AddItemToDb(req, client) {
		return &pb.AddItemResponse{Message: "Success"}, nil
	}
	defer cancel()
	return &pb.AddItemResponse{Message: "Error"}, nil
}

func (s *server) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	client, cancel := ConToDb()
	response := service.GetItem(client)
	defer cancel()
	return response, nil
}

func (s *server) DelItem(ctx context.Context, req *pb.DelItemRequest) (*pb.DelItemResponse, error) {
	client, cancel := ConToDb()
	if service.DelItem(client, req) {
		return &pb.DelItemResponse{Message: "Success"}, nil
	}
	defer cancel()
	return &pb.DelItemResponse{Message: "Error"}, nil
}

func (s *server) GetInventory(ctx context.Context, req *pb.GetInventoryRequest) (*pb.GetItemResponse, error) {
	client, cancel := ConToDb()
	fmt.Println(req)
	response := service.ItemsToProto(service.GetInventory(client, req))
	defer cancel()
	return response, nil
}

func ConToDb() (*mongo.Client, context.CancelFunc) {
	const connectionString = "mongodb://Stock:Stock@localhost:27021"
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	fmt.Println("Connected successfully to MongoDB")
	return client, cancel
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterInventoryServiceServer(s, &server{})
	reflection.Register(s)
	log.Println("Starting microservice on :50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

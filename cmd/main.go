package main

import (
	"CustomFileStorage/api/controllers"
	"CustomFileStorage/api/server"
	"CustomFileStorage/repositories/database/MongoDB"
	"CustomFileStorage/service"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func main() {
	router := gin.Default()

	mongoURI := os.Getenv("MONGO_URI")

	//mongoURI := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Disconnect(context.Background())
	mongo := client.Database("Storage")
	Queries := MongoDB.NewQueries(mongo)
	Service := service.NewService(Queries)
	Storage := controllers.NewStorageControllers(Service)
	s := server.NewServer(router, Storage)
	if err := s.Start(":9090"); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"context"
	"fmt"
	bookpb "gRPC_Mongo_HTTP_TestTask/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

type BookServiceServer struct {}

var db *mongo.Client
var bookdb *mongo.Collection
var mongoCtx context.Context

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50051...")

	//Прослушиваем порт
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	//инициализируем сервер
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	srv := &BookServiceServer{}
	bookpb.RegisterBookServiceServer(s, srv)

	//Подключаемся в МонгоБД
	fmt.Println("Connecting to MongoDB...")
	mongoCtx = context.Background()
	db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error for connecting to MongoDB: %v", err)
	}

	//Проверяем подключение к БД
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}

	bookdb = db.Database("mydb").Collection("book")

	//Запускаем сервер \ Отключаемся командой CTRL+C
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50051")

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt)

	<-c

	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Closing MongoDB connection")
	db.Disconnect(mongoCtx)
	fmt.Println("Done.")
}

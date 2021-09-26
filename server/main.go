package main

import (
	"context"
	"fmt"
	"gRPC_Mongo_HTTP_TestTask"
	"gRPC_Mongo_HTTP_TestTask/handler/handlerGRPC"
	"gRPC_Mongo_HTTP_TestTask/store"

	bookpb "gRPC_Mongo_HTTP_TestTask/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {

	bd, err := gRPC_Mongo_HTTP_TestTask.ConnectDatabase()
	if err != nil {
		log.Fatalf("Error for connecting to MongoDB: %v", err)
	}

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
	srv := &handlerGRPC.BookServiceServer{}

	bookpb.RegisterBookServiceServer(s, srv)

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
	gRPC_Mongo_HTTP_TestTask.DBClose()
	log.Println("closing MongoDB connection")
	fmt.Println("Done.")
}

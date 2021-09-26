package main

import (
	"fmt"
	"gRPC_Mongo_HTTP_TestTask"
	"gRPC_Mongo_HTTP_TestTask/handler/handlerGRPC"

	bookpb "gRPC_Mongo_HTTP_TestTask/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {

	//Подключаемся к БД
	gRPC_Mongo_HTTP_TestTask.ConnectDatabase()
	log.Println("Connecting to MongoDB!")

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

	s.Stop()
	fmt.Println("\nStopping the server...")

	err = listener.Close()
	if err != nil {
		log.Println("Error closing MongoDB connection")
	} else {
		fmt.Println("Closing MongoDB connection")
	}

	err = gRPC_Mongo_HTTP_TestTask.DBClose()
	if err != nil {
		log.Println("Error in closed Connection to MongoDB.")
	}
	log.Println("closing MongoDB connection")

	fmt.Println("Done.")
}

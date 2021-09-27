package main

import (
	"fmt"
	"gRPC_Mongo_HTTP_TestTask"
	"gRPC_Mongo_HTTP_TestTask/handler"
	bookpb "gRPC_Mongo_HTTP_TestTask/proto"
	"gRPC_Mongo_HTTP_TestTask/server/serverHTTP"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	gRPC_Mongo_HTTP_TestTask.ConnectDatabase()
	log.Println("Connecting to MongoDB!")

	//Прослушиваем порт
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	//инициализируем сервер gRPC
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	srv := &handler.BookServiceServer{}

	bookpb.RegisterBookServiceServer(s, srv)


	//Запускаем сервер gRPC \ Отключаемся командой CTRL+C
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50051")

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt)

	//Подключаем serverHTTP
	serverHTTP.NewServerHTTP()
	log.Println("Connecting serverHTTP")

	<-c

	fmt.Println("\nStopping the server...")
	s.Stop()

	//Закрываем прослушивание порта
	listener.Close()

	//Отсоединяемся от MongoDB
	log.Println("closing MongoDB connection")

	log.Println("Done.")
}

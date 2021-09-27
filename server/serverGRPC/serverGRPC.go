package serverGRPC

import (
	"flag"
	"fmt"
	"gRPC_Mongo_HTTP_TestTask/handler"
	bookpb "gRPC_Mongo_HTTP_TestTask/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)


func NewServerGRPC() {

	addr := flag.String("addr", ":4112", "Сетевой адрес сервера")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

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
	infoLog.Println("Запуск сервера gRPC на порту 50051")

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt)

	infoLog.Printf("Запуск сервера на %s", *addr)
	errorLog.Fatal(err)

	<-c

	s.Stop()
	fmt.Println("\nStopping the server...")

	//Закрываем прослушивание порта
	err = listener.Close()
	if err != nil {
		log.Println("Error closing MongoDB connection")
	} else {
		fmt.Println("Closing MongoDB connection")
	}

}

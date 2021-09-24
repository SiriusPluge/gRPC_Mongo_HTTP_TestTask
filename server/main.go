package main

import (
	"context"
	"fmt"

	bookpb "gRPC_Mongo_HTTP_TestTask/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
)

type BookServiceServer struct {
}

var db *mongo.Client
var bookdb *mongo.Collection
var mongoCtx context.Context

type BookItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Name     string             `bson:"name"`
	Tag      string             `bson:"tag"`
}

func (s *BookServiceServer) CreateBook(ctx context.Context, req *bookpb.CreateBookReq) (*bookpb.CreateBookRes, error) {
	//преобразования данных в BSON
	book := req.GetBook()
	data := BookItem{
		// ID:       primitive.NilObjectID,
		AuthorID: book.GetAuthorId(),
		Name:    book.GetName(),
		Tag:  book.GetTag(),
	}

	//запись данных в БД
	result, err := bookdb.InsertOne(mongoCtx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	book.Id = oid.Hex()

	return &bookpb.CreateBookRes{Book: book}, nil
}

func (s *BookServiceServer) ReadBook(ctx context.Context, req *bookpb.ReadBookReq) (*bookpb.ReadBookRes, error) {
	//конвертируем ID
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	// находим книгу по ID и записываем декодированную информацию
	result := bookdb.FindOne(ctx, bson.M{"_id": oid})
	data := BookItem{}
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find blog with Object Id %s: %v", req.GetId(), err))
	}

	// приводим к типу ответа
	response := &bookpb.ReadBookRes{
		Blog: &bookpb.Book{
			Id:       oid.Hex(),
			AuthorId: data.AuthorID,
			Name:    data.Name,
			Tag:  data.Tag,
		},
	}
	return response, nil
}

func (s *BookServiceServer) DeleteBook(ctx context.Context, req *bookpb.DeleteBookReq) (*bookpb.DeleteBookRes, error) {
	//конвертируем ID
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	// находим и удаляем книгу в БД
	_, err = bookdb.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete blog with id %s: %v", req.GetId(), err))
	}

	return &bookpb.DeleteBookRes{
		Success: true,
	}, nil
}

func (s *BookServiceServer) UpdateBook(ctx context.Context, req *bookpb.UpdateBookReq) (*bookpb.UpdateBookRes, error) {

	book := req.GetBook()

	// конвертируем ID
	oid, err := primitive.ObjectIDFromHex(book.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied blog id to a MongoDB ObjectId: %v", err),
		)
	}

	// конвертируем для БД
	update := bson.M{
		"author_id": book.GetAuthorId(),
		"name":      book.GetName(),
		"tag":    book.GetTag(),
	}

	// конвертируем id для БД
	filter := bson.M{"_id": oid}

	// возвращаем обнавленный документ БД
	result := bookdb.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	// декод для ответа
	decoded := BookItem{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find blog with supplied ID: %v", err),
		)
	}

	return &bookpb.UpdateBookRes{
		Book: &bookpb.Book{
			Id:       decoded.ID.Hex(),
			AuthorId: decoded.AuthorID,
			Name:    decoded.Name,
			Tag:  decoded.Tag,
		},
	}, nil
}

func (s *BookServiceServer) ListBook(req *bookpb.ListBookReq, stream bookpb.BookService_ListBookServer) error {

}

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

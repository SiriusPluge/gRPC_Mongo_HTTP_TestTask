package internalHTTP

import (
	"context"
	"fmt"
	"gRPC_Mongo_HTTP_TestTask"
	bookpb "gRPC_Mongo_HTTP_TestTask/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sync"
)

type BookItemCreate struct {
	ID       string `bson:"_id,omitempty"`
	AuthorID string `bson:"author_id"`
	Name     string `bson:"name"`
	Tag      string `bson:"tag"`
}

type BookStore struct {
	sync.Mutex
	Book map[string][]BookItemCreate
}

func New() *BookStore {
	us := &BookStore{}
	us.Book = make(map[string][]BookItemCreate)

	return us
}

var mongoCtx context.Context

func (s *Server) CreateBook(author_ID string, name string, tag string) string {
	s.Store.Lock()
	defer s.Store.Unlock()

	log.Println("Create book")
	var createBooks BookItemCreate
	//createBooks.ID = id
	createBooks.AuthorID = author_ID
	createBooks.Name = name
	createBooks.Tag = tag

	//запись данных в БД
	result, err := gRPC_Mongo_HTTP_TestTask.Collections.InsertOne(mongoCtx, createBooks)
	if err != nil {
		fmt.Sprintf("Internal error: %v", err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	createBooks.ID = oid.Hex()

	return createBooks.ID
}

func (s *Server) GetUser(id string) *bookpb.Book {
	log.Println("Get book")

	//конвертируем ID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Sprintf("Could not convert to ObjectId: %v", err)
	}

	// находим книгу по ID и записываем декодированную информацию
	result := gRPC_Mongo_HTTP_TestTask.Collections.FindOne(mongoCtx, bson.M{"_id": oid})
	data := &bookpb.Book{}
	if err := result.Decode(&data); err != nil {
		fmt.Sprintf("Could not find blog with Object Id %s: %v", id, err)
	}

	// приводим к типу ответа
	response := &bookpb.Book{
			Id:       oid.Hex(),
			AuthorId: data.AuthorId,
			Name:     data.Name,
			Tag:      data.Tag,
		}

	return response
}

func (s *Server) DeleteBook(id string) (*bookpb.DeleteBookRes, error) {
	log.Println("Delete book")
	//конвертируем ID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	// находим и удаляем книгу в БД
	_, err = gRPC_Mongo_HTTP_TestTask.Collections.DeleteOne(mongoCtx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete blog with id %s: %v", id, err))
	}

	return &bookpb.DeleteBookRes{
		Success: true,
	}, nil
}

func (s *Server) UpdateBook(id string, author_ID string, name string, tag string) (*bookpb.UpdateBookReq, error) {

	// конвертируем ID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
			fmt.Sprintf("Could not convert the supplied blog id to a MongoDB ObjectId: %v", err)
	}

	// конвертируем для БД
	update := bson.M{
		"author_id": author_ID,
		"name":      name,
		"tag":       tag,
	}

	// конвертируем id для БД
	filter := bson.M{"_id": oid}

	// возвращаем обнавленный документ БД
	result := gRPC_Mongo_HTTP_TestTask.Collections.FindOneAndUpdate(mongoCtx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	// декод для ответа
	decoded := &BookItemCreate{}
	err = result.Decode(&decoded)
	if err != nil {
			fmt.Sprintf("Could not find blog with supplied ID: %v", err)
	}

	return (*bookpb.UpdateBookReq)(&bookpb.UpdateBookRes{
		Book: &bookpb.Book{
			Id:       decoded.ID,
			AuthorId: decoded.AuthorID,
			Name:     decoded.Name,
			Tag:      decoded.Tag,
		},
	}), nil
}

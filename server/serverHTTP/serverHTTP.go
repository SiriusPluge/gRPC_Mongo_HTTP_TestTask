package serverHTTP

import (
	"gRPC_Mongo_HTTP_TestTask/server/internalHTTP"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewServerHTTP() {
	mux := mux.NewRouter()
	server := internalHTTP.NewUserServer()

	mux.HandleFunc("/api/book", server.CreateBookHandler).Methods("POST")
	mux.HandleFunc("/api/book/get", server.GetBookHandler).Methods("GET")
	mux.HandleFunc("/api/book/delete", server.DeleteBookHandler).Methods("DELETE")
	mux.HandleFunc("/api/book/put", server.UpdateBookHandler).Methods("PUT")

	log.Fatal(http.ListenAndServe("localhost:4112", mux))
}

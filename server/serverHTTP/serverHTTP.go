package serverHTTP

import (
	"flag"
	"gRPC_Mongo_HTTP_TestTask/server/internalHTTP"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func NewServerHTTP() {

	addr := flag.String("addr", ":4112", "Сетевой адрес сервера")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := mux.NewRouter()
	server := internalHTTP.NewUserServer()

	mux.HandleFunc("/api/book", server.CreateBookHandler).Methods("POST")
	mux.HandleFunc("/api/book/get", server.GetBookHandler).Methods("GET")
	mux.HandleFunc("/api/book/delete", server.DeleteBookHandler).Methods("DELETE")
	mux.HandleFunc("/api/book/put", server.UpdateBookHandler).Methods("PUT")

	infoLog.Printf("Запуск сервера на %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}

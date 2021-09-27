package internalHTTP

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
)

type BookItem struct {
	ID       string `json:"_id,omitempty"`
	AuthorID string `json:"authorID"`
	Name     string `json:"name"`
	Tag      string `json:"tag"`
}

func (s *Server) CreateBookHandler(w http.ResponseWriter, req *http.Request)  {

	log.Printf("handling book create at %s\n", req.URL.Path)

	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	var jsonData BookItem
	jsonDataFromHttp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)
	if err != nil {
		panic(err)
	}

	data := s.CreateBook(jsonData.AuthorID, jsonData.Name, jsonData.Tag)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *Server) GetBookHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get book at %s\n", req.URL.Path)

	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	var jsonData BookItem
	jsonDataFromHttp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)
	if err != nil {
		panic(err)
	}

	book := s.GetUser(jsonData.ID)

	js, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *Server) DeleteBookHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling delete book at %s\n", req.URL.Path)

	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	var jsonData BookItem
	jsonDataFromHttp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)
	if err != nil {
		panic(err)
	}

	book, _ := s.DeleteBook(jsonData.ID)

	js, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *Server) UpdateBookHandler(w http.ResponseWriter, req *http.Request) {

	log.Printf("handling book update at %s\n", req.URL.Path)

	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	var jsonData BookItem
	jsonDataFromHttp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)
	if err != nil {
		panic(err)
	}

	data, _ := s.UpdateBook(jsonData.ID, jsonData.AuthorID, jsonData.Name, jsonData.Tag)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

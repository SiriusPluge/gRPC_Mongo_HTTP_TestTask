package internalHTTP

type Server struct {
	Store *BookStore
}

func NewUserServer() *Server {
	Store := New()
	return &Server{Store: Store}
}

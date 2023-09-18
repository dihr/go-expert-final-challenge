package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type EndpointInfo struct {
	method string
	h      http.HandlerFunc
	path   string
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]EndpointInfo
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]EndpointInfo),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(name string, method string, path string, handler http.HandlerFunc) {
	s.Handlers[name] = EndpointInfo{
		method: method,
		h:      handler,
		path:   path,
	}
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		if handler.method == "GET" {
			s.Router.Get(handler.path, handler.h)
		}
		if handler.method == "POST" {
			s.Router.Post(handler.path, handler.h)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}

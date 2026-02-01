// package hello exposes the hello server for saying hello
package hello

import (
	"context"
	"net/http"
	"strconv"

	"example.com/htmx-test/render"
	"github.com/a-h/templ"
)

// Server - the server for giving you nice hello greetings
type Server struct {
	srv http.Server
}

// NewServer - creates a new server
func NewServer(port int) *Server {
	mux := http.NewServeMux()

	// serve one route on `/` which will be our hello page
	mux.HandleFunc("/", handleHelloPage)

	return &Server{
		srv: http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: mux,
		},
	}
}

func handleHelloPage(w http.ResponseWriter, r *http.Request) {
	var (
		name = ""
		opts []func(*templ.ComponentHandler)
	)

	q := r.URL.Query()

	if f := q.Get("fragment"); f != "" {
		opts = append(opts, templ.WithFragments(f))
	}

	name = r.FormValue("name")

	templ.Handler(render.HelloPage(name), opts...).ServeHTTP(w, r)
}

// Serve - starts the server
func (s *Server) Serve() error {
	return s.srv.ListenAndServe()
}

// Close - gracefully closes the server
func (s *Server) Close() error {
	return s.srv.Shutdown(context.Background())
}

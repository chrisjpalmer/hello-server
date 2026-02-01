package web

import (
	"context"
	"net/http"
	"strconv"

	"example.com/htmx-test/render"
	"github.com/a-h/templ"
)

type Server struct {
	srv http.Server
}

func NewServer(port int) *Server {
	mux := http.NewServeMux()

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

func (s *Server) Serve() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Close() error {
	return s.srv.Shutdown(context.Background())
}

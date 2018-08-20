package web

import (
	"net/http"
	"fmt"
	"context"
	"path"
	"github.com/davegarred/woodinville/domain"
)

const userCookie = "im"

var maps_key string

type Server struct {
	s            *http.Server
	serverClosed chan struct{}
}

type pathResolver map[string]func(http.ResponseWriter, *http.Request)

func Serve(key string) Server {
	maps_key = key
	srv := &http.Server{
		Addr:    ":8000",
		Handler: defaultPathResolver(),
	}
	serverClosed := make(chan struct{})
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println("Server shutdown gracefully")
		} else {
			fmt.Printf("Server crashed: %v\n", err)
		}
		close(serverClosed)
	}()
	return Server{srv, serverClosed}
}

func (srv Server) Shutdown() {
	srv.s.Shutdown(context.Background())
	<-srv.serverClosed
}

func (pr pathResolver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pr.resolve(r.Method + " " + r.URL.Path)(w, r)
}
func defaultPathResolver() pathResolver {
	pr := make(map[string]func(http.ResponseWriter, *http.Request))
	pr["GET /user"] = userFilter(userHandler)
	pr["GET /location"] = userFilter(locationHandler)
	pr["GET /details"] = userFilter(detailsHandler)
	pr["GET /do/*"] = userFilter(func(writer http.ResponseWriter, request *http.Request, id domain.UserId) {
		fmt.Println(request.URL)
	})
	return pr
}

func (pr pathResolver) resolve(u string) func(http.ResponseWriter, *http.Request) {
	for pattern, handler := range pr {
		if ok, err := path.Match(pattern, u); ok && err == nil {
			return corsFilter(handler)
		}
	}
	return corsFilter(roothandler)
}

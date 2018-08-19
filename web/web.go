package web

import (
	"net/http"
	"fmt"
	"context"
	"github.com/davegarred/woodinville/storage"
	"path"
)

const userCookie = "im"

type Server struct {
	s            *http.Server
	serverClosed chan struct{}
}

type pathResolver map[string]func(http.ResponseWriter, *http.Request)

func Serve() Server {
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

func roothandler(w http.ResponseWriter, r *http.Request) {
	userId := userIdFromParams(r)
	if len(userId) > 0 {
		http.SetCookie(w, &http.Cookie{
			Name:   userCookie,
			Value:  userId,
			MaxAge: 0,
		})
		fmt.Println("User " + userId + " logged in")
	}
}

func userIdFromParams(r *http.Request) string {
	signinValues := r.URL.Query()[userCookie]
	if len(signinValues) == 1 {
		return signinValues[0]
	}
	return ""
}

func userFilter(f func(http.ResponseWriter, *http.Request, storage.UserId)) func(http.ResponseWriter, *http.Request) {
	result := func(w http.ResponseWriter, r *http.Request) {
		userCode := userIdFromParams(r)
		if userCode == "" {
			cookie, err := r.Cookie(userCookie)
			if err != nil {
				if err != http.ErrNoCookie {
					panic(err)
				}
				w.WriteHeader(401)
				return
			}
			userCode = cookie.Value
		}
		userId := storage.UserId(userCode)
		user := storage.FindUser(userId)
		if user == nil {
			w.WriteHeader(403)
			return
		}
		f(w, r, userId)
	}
	return result
}
func corsFilter(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		f(w,r)
	}
}



func (pr pathResolver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pr.resolve(r.Method + " " + r.URL.Path)(w, r)
}
func defaultPathResolver() pathResolver {
	pr := make(map[string]func(http.ResponseWriter, *http.Request))
	pr["GET /user"] = userFilter(userHandler)
	pr["GET /location"] = userFilter(locationHandler)
	pr["GET /do/*"] = userFilter(func(writer http.ResponseWriter, request *http.Request, id storage.UserId) {
		fmt.Println(request.URL)
	})
	return pr
}

func (pr pathResolver) resolve(u string) func(http.ResponseWriter, *http.Request) {
	for pattern, handler := range pr {
		if ok, err := path.Match(pattern, u); ok &&  err == nil {
			return corsFilter(handler)
		}
	}
	return corsFilter(roothandler)
}

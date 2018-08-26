package web

import (
	"net/http"
	"github.com/davegarred/woodinville/storage"
	"github.com/davegarred/woodinville/domain"
)

func userFilter(f func(http.ResponseWriter, *http.Request, domain.UserId)) func(http.ResponseWriter, *http.Request) {
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
		userId := domain.UserId(userCode)
		user := storage.FindUser(userId)
		if user == nil {
			w.WriteHeader(403)
			return
		}
		f(w, r, userId)
	}
	return result
}

func userIdFromParams(r *http.Request) string {
	signinValues := r.URL.Query()[userCookie]
	if len(signinValues) == 1 {
		return signinValues[0]
	}
	return ""
}

func corsFilter(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		f(w,r)
	}
}


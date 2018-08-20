package web

import (
	"net/http"
	"github.com/davegarred/woodinville/storage"
	"encoding/json"
	"fmt"
	"github.com/davegarred/woodinville/domain"
)

func locationHandler(w http.ResponseWriter, r *http.Request, _ domain.UserId) {
	ser, err := json.Marshal(storage.FindArea())
	if err != nil {
		http.NotFound(w,r)
	}
	_,err = w.Write(ser)
	if err != nil {
		panic(err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request, userId domain.UserId) {
	ser, err := json.Marshal(storage.FindUser(userId))
	if err != nil {
		http.NotFound(w,r)
	}
	_,err = w.Write(ser)
	if err != nil {
		panic(err)
	}
}

type details struct {
	MapsApi string `json:"maps_api"`
}
func detailsHandler(w http.ResponseWriter, r *http.Request, userId domain.UserId) {
	ser, err := json.Marshal(details{maps_key})
	if err != nil {
		http.NotFound(w,r)
	}
	_,err = w.Write(ser)
	if err != nil {
		panic(err)
	}
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
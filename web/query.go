package web

import (
	"net/http"
	"github.com/davegarred/woodinville/storage"
	"encoding/json"
)

func locationHandler(w http.ResponseWriter, r *http.Request, _ storage.UserId) {
	ser, err := json.Marshal(storage.FindArea())
	if err != nil {
		http.NotFound(w,r)
	}
	_,err = w.Write(ser)
	if err != nil {
		panic(err)
	}
}
func userHandler(w http.ResponseWriter, r *http.Request, userId storage.UserId) {
	ser, err := json.Marshal(storage.FindUser(userId))
	if err != nil {
		http.NotFound(w,r)
	}
	_,err = w.Write(ser)
	if err != nil {
		panic(err)
	}
}
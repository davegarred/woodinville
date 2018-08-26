package web

import (
	"net/http"
	"github.com/davegarred/woodinville/storage"
	"encoding/json"
	"fmt"
	"github.com/davegarred/woodinville/domain"
	"io/ioutil"
)

func areaHandler(w http.ResponseWriter, r *http.Request, _ domain.UserId) {
	ser, err := json.Marshal(storage.FindArea())
	returnResult(w, r, ser, err)
}

func locationHandler(w http.ResponseWriter, r *http.Request, _ domain.UserId) {
	ser, err := json.Marshal(storage.FindWineries())
	returnResult(w, r, ser, err)
}

func userHandler(w http.ResponseWriter, r *http.Request, userId domain.UserId) {
	ser, err := json.Marshal(storage.FindUser(userId))
	returnResult(w, r, ser, err)
}

func returnResult(w http.ResponseWriter, r *http.Request, ser []byte, err error) {
	if err != nil {
		http.NotFound(w, r)
	}
	_, err = w.Write(ser)
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
		http.NotFound(w, r)
	}
	_, err = w.Write(ser)
	if err != nil {
		panic(err)
	}
}

func commandHandler(w http.ResponseWriter, r *http.Request, _ domain.UserId) {
	commandName := r.URL.Path[4:]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	switch commandName {
	case "CreateUser":
		command := domain.CreateUser{}
		err := json.Unmarshal(body, &command)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(400)
			return
		}
		storage.Dispatch(command)
	case "SetUserAdmin":
		command := &domain.SetUserAdmin{}
		err := json.Unmarshal(body, command)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(400)
			return
		}
		storage.Dispatch(*command)
	case "AddVisit":
		command := &domain.AddVisit{}
		err := json.Unmarshal(body, command)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(400)
			return
		}
		storage.Dispatch(*command)
	case "CreateWinery":
		command := &domain.CreateWinery{}
		err := json.Unmarshal(body, command)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(400)
			return
		}
		storage.Dispatch(*command)
	case "UpdateWineryPosition":
		command := &domain.UpdateWineryPosition{}
		err := json.Unmarshal(body, command)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(400)
			return
		}
		storage.Dispatch(*command)
	case "UpdateWineryAddress":
		command := &domain.UpdateWineryAddress{}
		err := json.Unmarshal(body, command)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(400)
			return
		}
		storage.Dispatch(*command)
	case "RecommendWinery":
		command := &domain.RecommendWinery{}
		err := json.Unmarshal(body, command)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(400)
			return
		}
		storage.Dispatch(*command)
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

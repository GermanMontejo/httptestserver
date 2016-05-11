package handlers

import (
	"encoding/json"
	"net/http"

	"errors"
	. "github.com/GermanMontejo/httptestserver/domain"
	"log"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUser")
	users, err := json.Marshal(UserStore)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateUser")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error:%v", err)
		return
	}
	err = validate(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error:%v", err)
		return
	}
	UserStore = append(UserStore, user)
	w.WriteHeader(http.StatusCreated)
	j, err := json.Marshal(UserStore)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error:%v", err)
		return
	}
	w.Write(j)
}

func validate(user User) error {
	for _, u := range UserStore {
		if u.Email == user.Email {
			return errors.New("This email already exists.")
		}
	}
	return nil
}

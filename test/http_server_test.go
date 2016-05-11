package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"encoding/json"
	"fmt"
	. "github.com/GermanMontejo/httptestserver/domain"
	. "github.com/GermanMontejo/httptestserver/handlers"
	"github.com/gorilla/mux"
)

func TestCreateUser(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", CreateUser).Methods("POST")
	// shorter way of directly creating a json payload
	userJson := `{"firstname":"german", "lastname":"montejo", "email":"gemontejo21@gmail.com"}`
	userReader := strings.NewReader(userJson)

	req, err := http.NewRequest("POST", "/users", userReader)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 201 {
		t.Errorf("HTTP status expected: 201, got:%d", w.Code)
	}
}

// sends http request using a multiplexor through ServeHTTP method
// ServeHTTP method executes the handler registered, in the matched route.
func TestGetUser(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("HTTP status expected: 200, got: %d", w.Code)
	}
	log.Println(w.Body.String())
}

func TestCreateUserClient(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", CreateUser).Methods("POST")
	server := httptest.NewServer(r)
	defer server.Close()
	usersUrl := fmt.Sprintf("%s/users", server.URL)
	fmt.Println(usersUrl)
	user := User{
		Firstname: "German",
		Lastname:  "Montejo",
		Email:     "gemontejo@gmail.com",
	}
	j, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	userReader := strings.NewReader(string(j))
	// userJson := `"firstname":"german", "lastname":"montejo", "email":"german_thegreat21@yahoo.com"`
	request, err := http.NewRequest("POST", usersUrl, userReader)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 201 {
		t.Errorf("HTTP status expected: 201, got:%d", res.StatusCode)
	}
}

// Sends an http request using a server. This uses the Do function, which accepts
// a request object and returns a response and an error object.
func TestGetUsersClient(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()
	usersUrl := fmt.Sprintf("%s/users", server.URL)
	request, err := http.NewRequest("GET", usersUrl, nil)
	if err != nil {
		t.Error(err)
	}
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("HTTP status expected: 200, got:%d", res.StatusCode)
	}
}

func TestUniqueEmail(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", CreateUser).Methods("POST")
	// longer way of creating a json payload. first we define a struct, then marshal
	// it, then it becomes our json payload. We then convert the slice of bytes to a string and then convert it to reader,
	// then pass it as an argument to NewRequest, which gives us a new Request object.
	userStruct := User{
		Firstname: "german",
		Lastname:  "montejo",
		Email:     "gemontejo@yahoo.com",
	}
	j, err := json.Marshal(userStruct)
	if err != nil {
		log.Println("Error:", err)
	}
	userReader := strings.NewReader(string(j))
	req, err := http.NewRequest("POST", "/users", userReader)
	if err != nil {
		log.Println("Error:", err)
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code == 400 {
		t.Error("Bad Request not expected, got:%d", w.Code)
	}
}

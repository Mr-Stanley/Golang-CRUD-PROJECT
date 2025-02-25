package handler

/*import (
	"GO-CRUD-PROJECT/internal/model"
	"GO-CRUD-PROJECT/internal/repository"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := model.User{
		Name:  "John Test",
		Email: "johntest@example.com",
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Error marshalling user: %v", err)
	}

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/users", CreateUser).Methods("POST")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status code %v, got %v", http.StatusCreated, status)
	}

	var createdUser model.User
	err = json.NewDecoder(rr.Body).Decode(&createdUser)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	if createdUser.Name != user.Name {
		t.Errorf("Expected user name %v, got %v", user.Name, createdUser.Name)
	}
	if createdUser.Email != user.Email {
		t.Errorf("Expected user email %v, got %v", user.Email, createdUser.Email)
	}
}

func TestGetUser(t *testing.T) {
	user := model.User{
		Name:  "Jane Test",
		Email: "janetest@example.com",
	}
	createdUser, err := repository.CreateUser(user)
	if err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	req, err := http.NewRequest("GET", "/users/"+strconv.Itoa(createdUser.ID), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/users/{id:[0-9]+}", GetUser).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}

	var fetchedUser model.User
	err = json.NewDecoder(rr.Body).Decode(&fetchedUser)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	if fetchedUser.ID != createdUser.ID {
		t.Errorf("Expected user ID %v, got %v", createdUser.ID, fetchedUser.ID)
	}
}
*/

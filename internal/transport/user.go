package transport

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	dtotransport "github.com/timur-danilchenko/project/internal/dto/transport"
	"github.com/timur-danilchenko/project/internal/service"
)

type UserTransport struct {
	Service *service.UserService
}

func (t *UserTransport) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userReq dtotransport.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		log.Printf("[ERROR] %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Can't decode JSON"))
		return
	}

	userRes, err := t.Service.CreateUser(r.Context(), userReq)
	if err != nil {
		log.Printf("[ERROR] %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userRes)
}

func (t *UserTransport) GetUserHandlerByID(w http.ResponseWriter, r *http.Request) {
	var userReq dtotransport.GetUserByIDRequest

	userReq.ID = uuid.MustParse(strings.TrimPrefix(r.URL.Path, "/"))

	userRes, err := t.Service.GetUserByID(r.Context(), userReq)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("No user with id{%s}", userReq.ID)))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userRes)
}

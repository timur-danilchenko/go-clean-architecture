package transport

import (
	"net/http"
)

func GreetingsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func SetUserTransport(router *http.ServeMux, userTransport *UserTransport) {
	userRouter := http.NewServeMux()

	userRouter.HandleFunc("/new", userTransport.CreateUserHandler)
	userRouter.HandleFunc("/{id}", userTransport.GetUserHandlerByID)

	router.Handle("/user/", http.StripPrefix("/user", userRouter))
}

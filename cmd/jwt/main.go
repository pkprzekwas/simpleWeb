package main

//https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b

import (
	"github.com/gorilla/mux"
	"github.com/pkprzekwas/simpleWeb/pkg/controllers"
	"github.com/pkprzekwas/simpleWeb/pkg/jwt"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.Use(jwt.JwtAuthentication)

	port := os.Getenv("PORT")
	if port != "" {
		port = "8080"
	}

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}

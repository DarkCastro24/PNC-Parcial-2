package main

import (
	"fmt"
	"log"
	"net/http"

	"PNC-Parcial-2/controllers"
	"PNC-Parcial-2/data"
	"PNC-Parcial-2/repositories"
	"PNC-Parcial-2/services"
	"github.com/gorilla/mux"
)

func main() {
	dsn := "user=postgres password=2002 dbname=parcial2 sslmode=disable"
	db := data.InitDB(dsn)
	fmt.Println("Conexion exitosa a PostgreSQL")

	repo := repositories.NewBookRepositoryGORM(db)
	service := services.NewUserService(repo)
	controller := controllers.BookController{Service: service}

	router := mux.NewRouter()
	controller.Routes(router)

	fmt.Println("Servidor corriendo en el puerto: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

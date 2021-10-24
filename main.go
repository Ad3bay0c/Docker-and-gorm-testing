package main

import (
	"fmt"
	"github.com/Ad3bay0c/dockersAndGormTesting/controllers"
	"github.com/Ad3bay0c/dockersAndGormTesting/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error: %v", err.Error())
	}
	router := gin.Default()
	DB := &db.PostgreSql{}
	DB.DbConnection()
	app := &controllers.Application{DB: DB}

	router.GET("/", app.Index)

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if PORT == ":" {
		PORT += "8080"
	}
	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Println("Working perfectly well")
	err = server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

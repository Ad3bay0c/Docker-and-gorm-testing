package main

import (
	"fmt"
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

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"Message": "Working Perfectly"})
	})
	PORT := fmt.Sprintf(":%s",os.Getenv("PORT"))
	if PORT == ":" {
		PORT += "8080"
	}
	server := &http.Server{
		Addr: PORT,
		Handler: router,
	}
	//db.DbConnection()
	log.Println("Working perfectly well")
	err = server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
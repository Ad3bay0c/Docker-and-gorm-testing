package main

import (
	"fmt"
	"github.com/Ad3bay0c/dockersAndGormTesting/controllers"
	"github.com/Ad3bay0c/dockersAndGormTesting/db"
	"github.com/Ad3bay0c/dockersAndGormTesting/models"
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
	DB.Db.AutoMigrate(&models.User{})

	//DB.Db.Create(&models.User{
	//	FirstName: "Adebayo",
	//	LastName:  "Adebayo",
	//	Author:    models.Author{Email: "aclonton", Password: "okay"},
	//	CreatedAt: time.Time{},
	//})
	//user := &models.User{}
	//DB.Db.Where("password = ?", "okay").Or("1 = 1").Limit(1).First(&user)
	//
	//user.FirstName = "Clinton"
	//DB.Db.Save(&user)
	//DB.Db.Model(&models.User{}).Where("id = ?", "1234resasd").Update("first_name", "Adebayo3")
	app := &controllers.Application{DB: DB}

	router.GET("/", app.Index)
	router.GET("/home", app.Index)

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

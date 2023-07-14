package main

import (
	"context"
	"log"
	"os"

	"baptiste.com/handlers"
	"baptiste.com/server"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	PROJECT_ID := os.Getenv("PROJECT_ID")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:      PORT,
		ProjectID: PROJECT_ID,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutesHome)
}

func BindRoutesHome(s server.Server, r *gin.Engine) {
	r.GET("/", handlers.HomeHandler)
	r.GET("/hello", handlers.HelloHandler)
	r.GET("/monthly-expenses/:id", handlers.GetMonthlyExpensesByIDHandler)
	r.POST("/monthly-expenses", handlers.PostMonthlyExpensesHandler)
	r.PATCH("/monthly-expenses", handlers.PatchMonthlyExpenseHandler)
}

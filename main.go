package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"baptiste.com/authenticator"
	"baptiste.com/handlers"
	"baptiste.com/server"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	fmt.Println("HOLAAAAA BIND ROUTES")
	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("auth-session", store))

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	r.GET("/", handlers.HomeHandler)
	r.GET("/hello", handlers.HelloHandler)

	//Rutas de MonthlyFixedExpenses
	r.POST("/monthly-fixed-expenses", handlers.PostMonthlyFixedExpense)
	r.GET("/monthly-fixed-expenses/:id", handlers.GetMonthlyFixedExpense)
	r.GET("/monthly-fixed-expenses", handlers.GetAllMonthlyFixedExpenses)
	r.PATCH("/monthly-fixed-expenses", handlers.UpdateMonthlyFixedExpense)

	//Rutas de Users
	r.POST("/users", handlers.PostUserHandler)
	r.GET("/users/:id", handlers.GetUser)
	r.GET("/users", handlers.GetAllUsers)
	r.PATCH("/users", handlers.UpdateUser)

	//Rutas de TrackingMonthlyFixedExpensesInsert
	r.POST("/tracking-monthly-fixed-expenses", handlers.PostTrackingMonthlyFixedExpense)
	r.GET("/tracking-monthly-fixed-expenses/:id", handlers.GetTrackingMonthlyFixedExpense)
	r.GET("/tracking-monthly-fixed-expenses", handlers.GetTrackingMonthlyFixedExpenses)
	r.PATCH("/tracking-monthly-fixed-expenses", handlers.UpdateTrackingMonthlyFixedExpense)

	//Login and Auth0
	r.GET("/login", handlers.HandlerLogin(auth))
	r.GET("/callback", handlers.HandlerCallback(auth))
	r.GET("/logout", handlers.HandlerLogOut)
}

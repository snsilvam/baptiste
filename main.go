package main

import (
	"context"
	"log"
	"os"

	"baptiste.com/config"
	"baptiste.com/handlers"
	"baptiste.com/middleware"
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
	AUTH0_AUDIENCE := os.Getenv("AUTH0_AUDIENCE")
	AUTH0_DOMAIN := os.Getenv("AUTH0_DOMAIN")
	CLIENT_ORIGIN_URL := os.Getenv("CLIENT_ORIGIN_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:          PORT,
		ProjectID:     PROJECT_ID,
		Domain:        AUTH0_DOMAIN,
		Audience:      AUTH0_AUDIENCE,
		SecureOptions: config.SecureOptions(),
		CorsOptions:   config.CorsOptions(CLIENT_ORIGIN_URL),
	})
	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutesHome)
}

func BindRoutesHome(s server.Server, r *gin.Engine) {
	r.GET("/", handlers.HomeHandler)
	r.GET("/hello", handlers.HelloHandler)

	//Routes of MonthlyFixedExpenses.
	r.POST("/monthly-fixed-expenses", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"create:monthlyFixed"}), handlers.PostMonthlyFixedExpense)
	r.GET("/monthly-fixed-expenses/:id", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"read:monthlyFixed"}), handlers.GetMonthlyFixedExpense)
	r.GET("/monthly-fixed-expenses", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"read:monthlyFixeds"}), handlers.GetAllMonthlyFixedExpenses)
	r.PATCH("/monthly-fixed-expenses", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"updated:monthlyFixed"}), handlers.UpdateMonthlyFixedExpense)

	//Routes of Users.
	r.POST("/users", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"create:user"}), handlers.PostUserHandler)
	r.GET("/users/:id", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"read:user"}), handlers.GetUser)
	r.GET("/users", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"read:users"}), handlers.GetAllUsers)
	r.PATCH("/users", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"updated:user"}), handlers.UpdateUser)

	//Routes of TrackingMonthlyFixedExpensesInsert.
	r.POST("/tracking-monthly-fixed-expenses", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"create:trackingMonthly"}), handlers.PostTrackingMonthlyFixedExpense)
	r.GET("/tracking-monthly-fixed-expenses/:id", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"read:trackingMonthly"}), handlers.GetTrackingMonthlyFixedExpense)
	r.GET("/tracking-monthly-fixed-expenses", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"read:trackingMonthlyFixeds"}), handlers.GetTrackingMonthlyFixedExpenses)
	r.PATCH("/tracking-monthly-fixed-expenses", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"updated:trackingMonthly"}), handlers.UpdateTrackingMonthlyFixedExpense)

	r.NoRoute(handlers.NotFoundHandler)
}

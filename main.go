package main

import (
	"context"
	"log"
	"os"

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
	AUTH0_AUDIENCE := os.Getenv("AUTH0_AUDIENCE")
	AUTH0_DOMAIN := os.Getenv("AUTH0_DOMAIN")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:     PORT,
		Domain:   AUTH0_DOMAIN,
		Audience: AUTH0_AUDIENCE,
	})
	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutesHome)
}

func BindRoutesHome(s server.Server, r *gin.Engine) {
	r.GET("/", handlers.HomeHandler)
	r.GET("/hello", handlers.HelloHandler)

	//Routes of Users.
	r.POST("/users", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"create:user"}), handlers.PostUserHandler)
	r.GET("/users/:id", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"read:user"}), handlers.GetUser)
	r.GET("/users", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"read:users"}), handlers.GetAllUsers)
	r.PATCH("/users", middleware.ValidateJWT(s.Config().Audience, s.Config().Domain),
		middleware.ValidatePermissions([]string{"updated:user"}), handlers.UpdateUser)
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"macadadi/stagetwo/db"
	"macadadi/stagetwo/repository"
	"macadadi/stagetwo/resources/user"
	"macadadi/stagetwo/services"

	"github.com/gin-gonic/gin"
)
const Port =":3001"
func main(){

   	db:= db.InitDB()
	defer db.Close()

	route := gin.New()
	route.Use(gin.Logger())



	userRepo := repository.NewUserRepository()

	userService := services.NewUserService(userRepo)

	user.Endpoints(route,db,userService)

	route.NoRoute(func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusNotFound,gin.H{"err_message":"end point not found try again later "})
	})

	server := http.Server{
		Addr: Port,
		Handler: route,
	}
	fmt.Printf("starting server on port %s \n", Port)
	log.Fatal(server.ListenAndServe())
}
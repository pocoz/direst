package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/dieqnt/direst/api"
	"github.com/dieqnt/direst/storage"
)

func main() {
	psql, err := storage.New()
	if err != nil {
		log.Fatalln(err)
	}

	user := api.New(psql)

	router := gin.Default()
	router.GET("/user/:id", user.UserGet)
	router.GET("/user", user.UserGetList)
	router.POST("/user", user.UserCreate)
	router.PUT("/user", user.UserUpdate)
	router.DELETE("/user/:id", user.UserDelete)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}

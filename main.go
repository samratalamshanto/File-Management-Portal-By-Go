package main

import (
	"File_Management_Portal/package/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	router.Register(r)
	log.Println("Application Run Successful")
	err := r.Run(":9090")
	if err != nil {
		log.Fatalf("Application Run Failed, {}", err)
	}

}

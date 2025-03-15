package main

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	dbErr := db.ConnectDB()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	r := gin.Default()

	err := r.Run(":9090")
	if err != nil {
		log.Fatal(err)
	}
}

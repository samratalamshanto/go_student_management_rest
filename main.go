package main

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/db"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/routes"
	"log"
)

func main() {
	dbErr := db.ConnectDB()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	r := routes.RegisterRoute()

	err := r.Run(":9090")
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/MnPutrav2/SIMRS-api/database"
	"github.com/MnPutrav2/SIMRS-api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	r := gin.Default()

	router.Patient(r, db)
	router.Payment(r, db)

	r.Run("localhost:8080")

}

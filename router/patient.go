package router

import (
	"database/sql"

	"github.com/MnPutrav2/SIMRS-api/controller"
	"github.com/gin-gonic/gin"
)

func Patient(r *gin.Engine, db *sql.DB) {

	controller.CreatePatient(r, db)
	controller.GetPatient(r, db)
	controller.GetAPI(r)

}

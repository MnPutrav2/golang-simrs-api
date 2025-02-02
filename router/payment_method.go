package router

import (
	"database/sql"

	"github.com/MnPutrav2/SIMRS-api/controller"
	"github.com/gin-gonic/gin"
)

func Payment(r *gin.Engine, db *sql.DB) {

	controller.GetPayment(r, db)

}

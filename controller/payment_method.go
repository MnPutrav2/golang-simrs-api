package controller

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/SIMRS-api/model"
	"github.com/gin-gonic/gin"
)

func GetPayment(r *gin.Engine, db *sql.DB) {
	r.GET("/payment", func(ctx *gin.Context) {
		id := ctx.Query("id")

		var payment model.PaymentMethod

		err := db.QueryRow("SELECT id, name FROM pay_method WHERE id = ?", id).Scan(&payment.ID, &payment.NamePay)
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, []model.PaymentMethod{
				{
					ID:      "0",
					NamePay: "0",
				},
			})
		} else {
			ctx.JSON(http.StatusOK, []model.PaymentMethod{
				{
					ID:      payment.ID,
					NamePay: payment.NamePay,
				},
			})
		}

	})
}

package controller

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/MnPutrav2/SIMRS-api/helper"
	"github.com/MnPutrav2/SIMRS-api/model"
	"github.com/gin-gonic/gin"
)

func CreatePatient(r *gin.Engine, db *sql.DB) {
	r.POST("/create", func(ctx *gin.Context) {
		var patient model.PatientData

		json.NewDecoder(ctx.Request.Body).Decode(&patient)

		db, err := db.Query(
			"INSERT INTO patient (medical_record, name, gender, birth_place, birth_date, district, village, subdistrict, city, province, postal_code, email, phone_number) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			patient.MedicalRecord,
			patient.Name,
			patient.Gender,
			patient.BirthPlace,
			patient.BirthDate,
			patient.District,
			patient.Village,
			patient.Subdistrict,
			patient.City,
			patient.Province,
			patient.PostalCode,
			patient.Email,
			patient.PhoneNumber,
		)

		helper.Error(err)

		defer db.Close()

		ctx.JSON(http.StatusCreated, []model.Response{
			{
				Status: "success",
				Data: []model.PatientData{
					{
						MedicalRecord: patient.MedicalRecord,
						Name:          patient.Name,
						Gender:        patient.Gender,
						BirthPlace:    patient.BirthPlace,
						BirthDate:     patient.BirthDate,
						District:      patient.District,
						Village:       patient.Village,
						Subdistrict:   patient.Subdistrict,
						City:          patient.City,
						Province:      patient.Province,
						PostalCode:    patient.PostalCode,
						Email:         patient.Email,
						PhoneNumber:   patient.PhoneNumber,
					},
				},
			},
		})

	})
}

func GetPatient(r *gin.Engine, db *sql.DB) {
	r.GET("/show", func(ctx *gin.Context) {
		mr := ctx.Query("mr")

		var patient model.PatientData
		err := db.QueryRow("SELECT medical_record, name, gender, birth_place, birth_date, district, village, subdistrict, city, province, postal_code, email, phone_number, registration  FROM patient WHERE medical_record = ?", mr).Scan(
			&patient.MedicalRecord,
			&patient.Name,
			&patient.Gender,
			&patient.BirthPlace,
			&patient.BirthDate,
			&patient.District,
			&patient.Village,
			&patient.Subdistrict,
			&patient.City,
			&patient.Province,
			&patient.PostalCode,
			&patient.Email,
			&patient.PhoneNumber,
			&patient.Registration,
		)
		helper.Error(err)

		ctx.JSON(http.StatusOK, []model.Response{
			{
				Status: "success",
				Data: []model.PatientData{
					{
						MedicalRecord: patient.MedicalRecord,
						Name:          patient.Name,
						Gender:        patient.Gender,
						BirthPlace:    patient.BirthPlace,
						BirthDate:     patient.BirthDate,
						District:      patient.District,
						Village:       patient.Village,
						Subdistrict:   patient.Subdistrict,
						City:          patient.City,
						Province:      patient.Province,
						PostalCode:    patient.PostalCode,
						Email:         patient.Email,
						PhoneNumber:   patient.PhoneNumber,
						Registration:  patient.Registration,
					},
				},
			},
		})
	})
}

func GetAPI(r *gin.Engine) {
	r.GET("/test", func(ctx *gin.Context) {
		response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
		helper.Error(err)

		responseData, err := ioutil.ReadAll(response.Body)
		helper.Error(err)

		json.NewDecoder(ctx.Request.Body).Decode(&responseData)

		ctx.JSON(http.StatusOK, string(responseData))
	})
}

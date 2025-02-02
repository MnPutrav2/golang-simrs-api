package model

type Response struct {
	Status string
	Data   []PatientData
}

type ResponseError struct {
	Status  string
	Message string
}

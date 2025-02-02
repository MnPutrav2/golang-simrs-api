package model

type PatientData struct {
	MedicalRecord int    `json:"medicalRecord"`
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	BirthPlace    string `json:"birthPlace"`
	BirthDate     string `json:"birthDate"`
	District      string `json:"district"`
	Village       string `json:"village"`
	Subdistrict   string `json:"subdistrict"`
	City          string `json:"city"`
	Province      string `json:"province"`
	PostalCode    int    `json:"postalCode"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phoneNumber"`
	Registration  string `json:"registration"`
}

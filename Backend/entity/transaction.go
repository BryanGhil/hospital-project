package entity

import "time"

type Transaction struct {
	ID         int
	PatientID  int
	MedicineID int
	Quantity   int
	TotalPrice float64
	Date       time.Time
	DoctorID   int
}

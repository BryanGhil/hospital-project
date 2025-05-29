package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type MedicalHistory struct {
	ID            int
	PatientID     int
	DoctorID      int
	Diagnosis     string
	Notes         string
	VisitDate     time.Time
	Prescriptions []Prescription
}

type Prescription struct {
	ID           int
	HistoryID    int
	MedicineID   int
	MedicineName string
	Dosage       string
	Quantity     int
	UnitPrice    decimal.Decimal
	TotalPrice   decimal.Decimal
}

package entity

import "time"

type Patient struct {
	ID        int
	FullName  string
	DOB       time.Time
	Gender    string
	Address   string
	Phone     string
	CreatedBy int
}

type ReqAddPatient struct {
	FullName  string
	DOB       time.Time
	Gender    string
	Address   string
	Phone     string
	CreatedBy int
}

type ReqUpdatePatient struct {
	FullName  *string
	DOB       *time.Time
	Gender    *string
	Address   *string
	Phone     *string
}
package repository

import (
	"backend/customerrors"
	"backend/entity"
	"context"
	"database/sql"
)

type PatientRepoItf interface {
	AddPatient(context.Context, entity.ReqAddPatient)(*entity.Patient, error)
}

type PatientRepoImpl struct {
}

func NewPatientRepo() PatientRepoImpl{
	return PatientRepoImpl{}
}

func (pr PatientRepoImpl) AddPatient(ctx context.Context, req entity.ReqAddPatient)(*entity.Patient, error) {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, customerrors.NewError(customerrors.DatabaseError, "internal server error")
	}

	var patient entity.Patient

	q := `
		insert into 
			patients (full_name, dob, gender, address, phone, created_by, created_at, updated_at)
		values 
			($1, $2, $3, $4, $5, $6, NOW(), NOW())
		returning 
			id, full_name, dob, gender, address, phone, created_by`

	err := tx.QueryRowContext(ctx, q, req.FullName, req.DOB, req.Gender, req.Address, req.Phone, req.CreatedBy).Scan(&patient.ID, &patient.FullName, &patient.DOB, &patient.Gender, &patient.Address, &patient.Phone, &patient.CreatedBy)

	if err != nil {
		return nil, customerrors.NewError(customerrors.DatabaseError, "error register user")
	}

	return &patient, nil
}
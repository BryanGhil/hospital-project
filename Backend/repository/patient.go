package repository

import (
	"backend/customerrors"
	"backend/entity"
	"context"
	"database/sql"
)

type PatientRepoItf interface {
	AddPatient(context.Context, entity.ReqAddPatient) (*entity.Patient, error)
	GetAllPatients(context.Context, entity.DefaultPageFilter) ([]entity.Patient, error)
	GetCountOfPatients(context.Context)(*int, error)
}

type PatientRepoImpl struct {
}

func NewPatientRepo() PatientRepoImpl {
	return PatientRepoImpl{}
}

func (pr PatientRepoImpl) AddPatient(ctx context.Context, req entity.ReqAddPatient) (*entity.Patient, error) {
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
		return nil, customerrors.NewError(customerrors.DatabaseError, "error occured")
	}

	return &patient, nil
}

func (pr PatientRepoImpl) GetAllPatients(ctx context.Context, filter entity.DefaultPageFilter) ([]entity.Patient, error) {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, customerrors.NewError(customerrors.DatabaseError, "internal server error")
	}

	q := `
		select  
			id, full_name, dob, gender, address, phone
		from 
			patients
		limit 
			$1
		offset 
			$2`

	rows, err := tx.QueryContext(ctx, q, filter.Limit, filter.Offset)
	if err != nil {
		return nil, customerrors.NewError(customerrors.DatabaseError, "internal server error")
	}

	defer rows.Close()

	var patients []entity.Patient

	for rows.Next() {
		var p entity.Patient
		if err := rows.Scan(&p.ID, &p.FullName, &p.DOB, &p.Gender, &p.Address, &p.Phone); err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return patients, nil
}

func (pr PatientRepoImpl) GetCountOfPatients(ctx context.Context)(*int, error) {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, customerrors.NewError(customerrors.DatabaseError, "internal server error")
	}

	var count int

	q := `
		select count(*)
		from patients`

	err := tx.QueryRowContext(ctx, q).Scan(&count)
	if err != nil {
		return nil, customerrors.NewError(customerrors.DatabaseError, "error occured")
	}
	return &count, nil

}
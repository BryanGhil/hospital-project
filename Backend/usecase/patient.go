package usecase

import (
	"backend/constant"
	"backend/customerrors"
	"backend/entity"
	"backend/repository"
	"context"
)

type PatientUsecaseItf interface {
	AddPatient(context.Context, entity.ReqAddPatient) (*entity.Patient, error)
	GetAllPatients(context.Context, entity.DefaultPageFilter) (*entity.GetPageResponse, error)
}

type PatientUsecaseImpl struct {
	pr  repository.PatientRepoItf
	trx repository.Transactor
}

func NewPatientUsecaseImpl(pr repository.PatientRepoItf, trx repository.Transactor) PatientUsecaseImpl {
	return PatientUsecaseImpl{
		pr:  pr,
		trx: trx,
	}
}

func (puc PatientUsecaseImpl) AddPatient(ctx context.Context, req entity.ReqAddPatient) (*entity.Patient, error) {
	data, err := puc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		user, ok := ctx.Value(constant.RequestUserId).(int)
		if !ok {
			return nil, customerrors.NewError(customerrors.DatabaseError, "internal server error")
		}

		req.CreatedBy = user

		res, err := puc.pr.AddPatient(ctx, req)
		if err != nil {
			return nil, err
		}

		return res, nil
	})
	if err != nil {
		return nil, err
	}

	patient, ok := data.(*entity.Patient)
	if !ok {
		return nil, customerrors.NewError(customerrors.CommonErr, "error occured")
	}

	return patient, nil
}

func (puc PatientUsecaseImpl) GetAllPatients(ctx context.Context, filter entity.DefaultPageFilter)(*entity.GetPageResponse, error) {
	data, err := puc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		patients, err := puc.pr.GetAllPatients(ctx, filter)
		if err != nil {
			return nil, err
		}

		pageCount, err := puc.pr.GetCountOfPatients(ctx)
		if err != nil {
			return nil, err
		}

		res := entity.GetPageResponse{Page: filter.Page, Limit: filter.Limit, CountData: *pageCount, Data: patients}

		return res, nil
	})
	if err != nil {
		return nil, err
	}

	res, ok := data.(entity.GetPageResponse)
	if !ok {
		return nil, customerrors.NewError(customerrors.CommonErr, "error occured")
	}

	return &res, nil
}

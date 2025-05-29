package usecase

import (
	"backend/customerrors"
	"backend/entity"
	"backend/repository"
	"context"
)

type MedicineUsecaseItf interface {
	AddMedicine(context.Context, entity.ReqAddMedicine) (*entity.Medicine, error)
}

type MedicineUsecaseImpl struct {
	mr  repository.MedicineRepoItf
	trx repository.Transactor
}

func NewMedicineUsecaseImpl(mr repository.MedicineRepoItf, trx repository.Transactor) MedicineUsecaseImpl {
	return MedicineUsecaseImpl{
		mr:  mr,
		trx: trx,
	}
}

func (muc MedicineUsecaseImpl) AddMedicine(ctx context.Context, req entity.ReqAddMedicine) (*entity.Medicine, error) {
	data, err := muc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {

		res, err := muc.mr.AddMedicine(ctx, req)
		if err != nil {
			return nil, err
		}

		return res, nil
	})
	if err != nil {
		return nil, err
	}

	medicine, ok := data.(*entity.Medicine)
	if !ok {
		return nil, customerrors.NewError(customerrors.CommonErr, "error occured")
	}

	return medicine, nil
}

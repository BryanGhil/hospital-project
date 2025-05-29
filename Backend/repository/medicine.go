package repository

import (
	"backend/customerrors"
	"backend/entity"
	"context"
	"database/sql"
)

type MedicineRepoItf interface {
	AddMedicine(context.Context, entity.ReqAddMedicine)(*entity.Medicine, error)
}

type MedicineRepoImpl struct {
}

func NewMedicineRepo() MedicineRepoImpl{
	return MedicineRepoImpl{}
}

func (mr MedicineRepoImpl) AddMedicine(ctx context.Context, req entity.ReqAddMedicine)(*entity.Medicine, error) {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, customerrors.NewError(customerrors.DatabaseError, "internal server error")
	}

	var medicine entity.Medicine

	q := `
		insert into 
			medicines (name, stock, price, created_at, updated_at)
		values 
			($1, $2, $3, NOW(), NOW())
		returning 
			id, name, stock, price`

	err := tx.QueryRowContext(ctx, q, req.Name, req.Stock, req.Price).Scan(&medicine.ID, &medicine.Name, &medicine.Stock, &medicine.Price)

	if err != nil {
		return nil, customerrors.NewError(customerrors.DatabaseError, "error occured")
	}

	return &medicine, nil
}
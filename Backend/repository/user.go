package repository

import (
	"backend/customerrors"
	"backend/entity"
	"context"
	"database/sql"
	"errors"
)

type UserRepoItf interface {
	CheckIsEmailExist(context.Context, entity.ReqRegisterUser) (bool, error)
	RegisterUser(context.Context, entity.ReqRegisterUser) (*entity.User, error)
	GetUserByEmail(context.Context, entity.ReqLoginUser) (*entity.User, error)
}

type UserRepoImpl struct {
}

func NewUserRepo() UserRepoImpl {
	return UserRepoImpl{}
}

func (ur UserRepoImpl) CheckIsEmailExist(ctx context.Context, req entity.ReqRegisterUser) (bool, error) {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return false, customerrors.NewError(customerrors.DatabaseError, "internal server error")
	}

	q := `
		select 
			id 
		from 
			users
		where 
			email = $1`

	err := tx.QueryRowContext(ctx, q, req.Email).Scan(new(string))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, customerrors.NewError(customerrors.DatabaseError, "internal server error")
	}

	return true, nil
}

func (ur UserRepoImpl) RegisterUser(ctx context.Context, req entity.ReqRegisterUser) (*entity.User, error) {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, customerrors.NewError(customerrors.DatabaseError, "internal server error")
	}

	var user entity.User

	q := `
		insert into 
			users (email, password, role_id, created_at, updated_at) 
		values 
			($1, $2, $3, NOW(), NOW())
		returning 
			id, email, role_id`

	err := tx.QueryRowContext(ctx, q, req.Email, req.Password, req.RoleId).Scan(&user.Id, &user.Email, &user.RoleId)

	if err != nil {
		return nil, customerrors.NewError(customerrors.DatabaseError, "error register user")
	}

	return &user, nil
}

func (ur UserRepoImpl) GetUserByEmail(ctx context.Context, req entity.ReqLoginUser) (*entity.User, error) {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, customerrors.NewError(customerrors.DatabaseError, "internal server error")
	}

	var user entity.User

	q := `
		select 
			id, email, password, role_id
		from 
			users
		where 
			email = $1`

	err := tx.QueryRowContext(ctx, q, req.Email).Scan(&user.Id, &user.Email, &user.Password, &user.RoleId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customerrors.NewError(customerrors.InvalidAction, "invalid credentials")
		}
		return nil, customerrors.NewError(customerrors.DatabaseError, "internal server error")
	}

	return &user, nil
}

package usecase

import (
	"backend/customerrors"
	"backend/entity"
	"backend/repository"
	"backend/util"
	"context"
)

type UserUsecaseItf interface {
	RegisterUser(context.Context, entity.ReqRegisterUser) (*entity.User, error)
	LoginUser(context.Context, entity.ReqLoginUser) (*entity.Token, error)
}

type UserUsecaseImpl struct {
	ur  repository.UserRepoItf
	trx repository.Transactor
}

func NewUserUsecaseImpl(ur repository.UserRepoItf, trx repository.Transactor) UserUsecaseImpl {
	return UserUsecaseImpl{
		ur:  ur,
		trx: trx,
	}
}

func (uuc UserUsecaseImpl) RegisterUser(ctx context.Context, req entity.ReqRegisterUser) (*entity.User, error) {
	data, err := uuc.trx.WithinTransaction(ctx, func(ctx context.Context) (any,error) {
		exist, err := uuc.ur.CheckIsEmailExist(ctx, req)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, customerrors.NewError(customerrors.AlreadyExist, "email already used")
		}

		hashPwd, err := util.HashPassword(req.Password)
		if err != nil {
			return nil, customerrors.NewError(customerrors.CommonErr, "error occured")
		}

		req.Password = hashPwd

		res, err := uuc.ur.RegisterUser(ctx, req)
		if err != nil {
			return nil,err
		}


		return res, nil
	})
	if err != nil {
		return nil, err
	}

	user, ok := data.(*entity.User)
	if !ok {
		return nil, customerrors.NewError(customerrors.CommonErr, "error occured")
	}

	return user, nil
}

func (uuc UserUsecaseImpl) LoginUser(ctx context.Context, req entity.ReqLoginUser) (*entity.Token, error) {
	data, err := uuc.trx.WithinTransaction(ctx, func(ctx context.Context) (any,error) {
		user, err := uuc.ur.GetUserByEmail(ctx, req)
		if err != nil {
			return nil,err
		}

		ok := util.CompareHashPassword(req.Password, user.Password)
		if !ok {
			return nil,customerrors.NewError(customerrors.InvalidAction, "invalid credentials")
		}

		tokenRes, err := util.GenerateJWTToken(user.Id, user.RoleId)
		if err != nil {
			return nil, customerrors.NewError(customerrors.CommonErr, "error login user")
		}


		return tokenRes,nil
	})
	if err != nil {
		return nil, err
	}
	
	token, ok := data.(string)
	if !ok {
		return nil, customerrors.NewError(customerrors.CommonErr, "error occured")
	}

	return &entity.Token{Token: token}, nil
}

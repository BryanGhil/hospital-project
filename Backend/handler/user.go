package handler

import (
	"backend/dto"
	"backend/entity"
	"backend/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uuc usecase.UserUsecaseItf
}

func NewUserHandler(uuc usecase.UserUsecaseItf) UserHandler {
	return UserHandler{
		uuc: uuc,
	}
}

func (uh UserHandler) RegisterUser(ctx *gin.Context) {
	var reqBody dto.ReqRegisterUser

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.Error(err)
		return
	}

	reqBodyEntity := entity.ReqRegisterUser{
		Email: reqBody.Email,
		Password: reqBody.Password,
		RoleId: reqBody.RoleId,
	}

	res, err := uh.uuc.RegisterUser(ctx, reqBodyEntity)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := dto.User{
		Id: res.Id,
		Email: res.Email,
		RoleId: res.RoleId,
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "successfully register user",
		Error: nil,
		Data: resDto,
	})
}

func (uh UserHandler) LoginUser(ctx *gin.Context) {
	var reqBody dto.ReqLoginUser

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.Error(err)
		return
	}

	reqBodyEntity := entity.ReqLoginUser{
		Email: reqBody.Email,
		Password: reqBody.Password,
	}

	res, err := uh.uuc.LoginUser(ctx, reqBodyEntity)
	if err != nil {
		ctx.Error(err)
		return 
	}

	resToken := dto.Token{
		Token: res.Token,
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "successfully login",
		Error: nil,
		Data: resToken,
	})
}
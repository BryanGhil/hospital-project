package handler

import (
	"backend/customerrors"
	"backend/dto"
	"backend/entity"
	"backend/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	puc usecase.PatientUsecaseItf
}

func NewPatientHandler(puc usecase.PatientUsecaseItf) PatientHandler {
	return PatientHandler{
		puc: puc,
	}
}

func (ph PatientHandler) AddPatient(ctx *gin.Context) {
	var reqBody dto.ReqAddPatient

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.Error(err)
		return
	}

	dob, err := time.Parse("2006-01-02", reqBody.DOB)
	if err != nil {
		ctx.Error(customerrors.NewError(customerrors.InvalidAction, "dob format not valid"))
		return
	}

	reqBodyEntity := entity.ReqAddPatient{
		FullName: reqBody.FullName,
		DOB: dob,
		Gender: reqBody.Gender,
		Address: reqBody.Address,
		Phone: reqBody.Phone,
	}

	res, err := ph.puc.AddPatient(ctx, reqBodyEntity)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := dto.Patient{
		ID: res.ID,
		FullName: res.FullName,
		DOB: res.DOB,
		Gender: res.Gender,
		Address: res.Address,
		Phone: res.Phone,
		CreatedBy: res.CreatedBy,
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "successfully add patient",
		Error: nil,
		Data: resDto,
	})
}
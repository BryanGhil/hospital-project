package handler

import (
	"backend/customerrors"
	"backend/dto"
	"backend/entity"
	"backend/usecase"
	"net/http"
	"strconv"
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
		DOB:      dob,
		Gender:   reqBody.Gender,
		Address:  reqBody.Address,
		Phone:    reqBody.Phone,
	}

	res, err := ph.puc.AddPatient(ctx, reqBodyEntity)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := dto.Patient{
		ID:        res.ID,
		FullName:  res.FullName,
		DOB:       res.DOB.Format("2006-01-02"),
		Gender:    res.Gender,
		Address:   res.Address,
		Phone:     res.Phone,
		CreatedBy: res.CreatedBy,
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "successfully add patient",
		Error:   nil,
		Data:    resDto,
	})
}

func (ph PatientHandler) GetAllPatients(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		ctx.Error(customerrors.NewError(customerrors.InvalidAction, "page not valid"))
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		ctx.Error(customerrors.NewError(customerrors.InvalidAction, "limit not valid"))
		return
	}

	offset := (pageInt - 1) * limitInt

	filter := entity.DefaultPageFilter{Page: pageInt, Limit: limitInt, Offset: offset}

	res, err := ph.puc.GetAllPatients(ctx, filter)
	if err != nil {
		ctx.Error(err)
		return
	}

	patients, ok := res.Data.([]entity.Patient)
	if !ok {
		ctx.Error(customerrors.NewError(customerrors.CommonErr, "error occured"))
		return
	}

	var resPatients []dto.Patient

	for _, val := range patients {
		resPatients = append(resPatients, dto.Patient{
			ID:       val.ID,
			FullName: val.FullName,
			DOB:      val.DOB.Format("2006-01-02"),
			Gender:   val.Gender,
			Address:  val.Address,
			Phone:    val.Phone,
		})
	}

	resDto := dto.GetPageResponse{
		Page:      res.Page,
		Limit:     res.Limit,
		CountData: res.CountData,
		Data:      resPatients,
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "successfully get all patients",
		Error:   nil,
		Data:    resDto,
	})
}

func (ph PatientHandler) GetPatientById(ctx *gin.Context) {
	id := ctx.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(customerrors.NewError(customerrors.InvalidAction, "id not valid"))
		return
	}

	res, err := ph.puc.GetPatientById(ctx, idInt)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := dto.Patient{
		ID: res.ID,
		FullName: res.FullName,
		DOB: res.DOB.Format("2006-01-02"),
		Gender: res.Gender,
		Address: res.Address,
		Phone: res.Phone,
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "successfully get patient data",
		Error:   nil,
		Data:    resDto,
	})
}

func (ph PatientHandler) UpdatePatients(ctx *gin.Context) {
	id := ctx.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(customerrors.NewError(customerrors.InvalidAction, "id not valid"))
		return
	}

	var reqBody dto.ReqUpdatePatient

	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.Error(err)
		return
	}

	var dob *time.Time = nil
	if reqBody.DOB != nil {
		dobDate, err := time.Parse("2006-01-02", *reqBody.DOB)
		if err != nil {
			ctx.Error(customerrors.NewError(customerrors.InvalidAction, "dob format not valid"))
			return
		}
		dob = &dobDate
	}

	reqBodyEntity := entity.ReqUpdatePatient{
		FullName: reqBody.FullName,
		DOB:      dob,
		Gender:   reqBody.Gender,
		Address:  reqBody.Address,
		Phone:    reqBody.Phone,
	}

	err = ph.puc.UpdatePatient(ctx, idInt, reqBodyEntity)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "successfully update patient",
		Error:   nil,
		Data:    nil,
	})
}

func (ph PatientHandler) DeletePatient(ctx *gin.Context) {
	id := ctx.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(customerrors.NewError(customerrors.InvalidAction, "id not valid"))
		return
	}

	err = ph.puc.DeletePatient(ctx, idInt)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "successfully delete patient",
		Error:   nil,
		Data:    nil,
	})
}

func (ph PatientHandler) RestoreDeletedPatient(ctx *gin.Context) {
	id := ctx.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(customerrors.NewError(customerrors.InvalidAction, "id not valid"))
		return
	}

	err = ph.puc.RestoreDeletedPatient(ctx, idInt)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "successfully restore deleted patient",
		Error:   nil,
		Data:    nil,
	})
}
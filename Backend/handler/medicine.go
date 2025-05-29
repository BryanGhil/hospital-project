package handler

import (
	"backend/dto"
	"backend/entity"
	"backend/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MedicineHandler struct {
	muc usecase.MedicineUsecaseItf
}

func NewMedicineHandler(muc usecase.MedicineUsecaseItf) MedicineHandler {
	return MedicineHandler{
		muc: muc,
	}
}

func (ph MedicineHandler) AddMedicine(ctx *gin.Context) {
	var reqBody dto.ReqAddMedicine

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.Error(err)
		return
	}

	reqBodyEntity := entity.ReqAddMedicine{
		Name: reqBody.Name,
		Stock: reqBody.Stock,
		Price: reqBody.Price,
	}

	res, err := ph.muc.AddMedicine(ctx, reqBodyEntity)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := dto.Medicine{
		ID: res.ID,
		Name: res.Name,
		Stock: res.Stock,
		Price: res.Price,
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "successfully add medicine",
		Error: nil,
		Data: resDto,
	})
}
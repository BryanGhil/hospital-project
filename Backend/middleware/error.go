package middleware

import (
	"backend/customerrors"
	"backend/dto"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		err := ctx.Errors[0]

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			validationErrors := make([]dto.ErrorRes,0)

			for _, fe := range ve {
				validationErrors = append(validationErrors, dto.ErrorRes{
					Field: fe.Field(),
					Message: fmt.Sprintf("invalid input on field %s", fe.Field()),
				})
			}

			ctx.JSON(http.StatusBadRequest,dto.Response{
				Success: false,
				Message: "invalid input",
				Error: validationErrors,
				Data: nil,
			})
			return
		}

		var ce *customerrors.CustomError
		if errors.As(err, &ce) {
			ctx.JSON(ce.GetHTTPErrorCode(), dto.Response{
				Success: false,
				Message: "error occured",
				Error: []dto.ErrorRes{{Message: ce.Error()}},
				Data: nil,
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Success: false,
				Message: "error occured",
				Error: []dto.ErrorRes{{Message: "internal server error"}},
				Data: nil,
		})
	}
}
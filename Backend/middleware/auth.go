package middleware

import (
	"backend/constant"
	"backend/customerrors"
	"backend/util"
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	authorizationHeaderKey = "authorization"
	authorizationTypeBearer = "bearer"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authHeader) == 0 {
			ctx.Error(customerrors.NewError(customerrors.Unauthorized, "authorization header not found"))
			ctx.Abort()
			return
		}

		splitHeader := strings.Split(ctx.GetHeader(authorizationHeaderKey), " ")
		if len(splitHeader) != 2 {
			ctx.Error(customerrors.NewError(customerrors.Unauthorized, "invalid token format"))
			ctx.Abort()
			return
		}

		authType := strings.ToLower(splitHeader[0])
		if authType != authorizationTypeBearer {
			ctx.Error(customerrors.NewError(customerrors.Unauthorized, "unsupported authorization type"))
			ctx.Abort()
			return
		}

		token := splitHeader[1]

		claims, err := util.ParseJWTToken(token)
		if err != nil {
			errMsg := "cannot parse token"
			if errors.Is(err, jwt.ErrTokenExpired){
				errMsg = "token has expired"
			}
			ctx.Error(customerrors.NewError(customerrors.Unauthorized, errMsg))
			ctx.Abort()
			return
		}

		userId, err := claims.GetSubject()
		if err != nil {
			ctx.Error(customerrors.NewError(customerrors.Unauthorized, "auth error"))
			ctx.Abort()
			return
		}

		intUserId, err := strconv.Atoi(userId)
		if err != nil {
			ctx.Error(customerrors.NewError(customerrors.CommonErr, "auth error"))
			ctx.Abort()
			return
		}

		ctx.Set(constant.RequestUserId, intUserId)
		ctx.Set(constant.RequestRoleId, claims.Role)
		ctx.Next()
	}
}

func AdminOnly() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		roleId, ok := ctx.Value(constant.RequestRoleId).(int)
		if !ok {
			ctx.Error(customerrors.NewError(customerrors.Unauthorized, "error auth"))
			ctx.Abort()
			return
		}

		if roleId != constant.RoleAdmin {
			ctx.Error(customerrors.NewError(customerrors.Unauthorized, "Unauthorized"))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
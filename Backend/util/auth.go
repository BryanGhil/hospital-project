package util

import (
	"backend/constant"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type customClaims struct {
	Role int `json:"role"`
	jwt.RegisteredClaims
}

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func CompareHashPassword(pwd string, hashPwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd)) == nil
}

func GenerateJWTToken(userId int, roleId int) (string, error) {
	secret := os.Getenv(constant.JWTSecret)

	now := time.Now()

	registeredClaim := customClaims{
		Role: roleId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:  os.Getenv(constant.AppName),
			Subject: strconv.Itoa(userId),
			IssuedAt: &jwt.NumericDate{
				Time: now,
			},
			ExpiresAt: &jwt.NumericDate{
				Time: now.Add(24 * time.Hour),
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, registeredClaim)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWTToken(tokenInp string) (*customClaims, error) {
	secret := os.Getenv(constant.JWTSecret)

	var myCustomClaim customClaims

	token, err := jwt.ParseWithClaims(tokenInp, &myCustomClaim, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired){
			return nil, fmt.Errorf("token expired")
		}
		return nil, err
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

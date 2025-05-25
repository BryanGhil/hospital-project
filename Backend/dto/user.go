package dto

type ReqRegisterUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleId   int    `json:"role_id" binding:"required"`
}

type User struct {
	Id     int    `json:"user_id"`
	Email  string `json:"email"`
	RoleId int    `json:"role_id"`
}

type ReqLoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	Token string `json:"token"`
}

package models

type UserLoginReq struct {
	Username string //用户名
	Password string //密码
}

type UserLoginRes struct {
	UserId   int    //用户ID
	Username string //用户名
	Phone    string //手机号
	Email    string //邮箱
}

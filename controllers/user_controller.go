package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

/*用户*/
type UserController interface {
	Regis(g *gin.Context)
	Login(g *gin.Context)
	Logout(g *gin.Context)
	UpdateInfo(g *gin.Context)
}

/*用户服务结构体*/
type UserService struct {
	SqlDB *sql.DB
}

type User struct {
	Id       int
	Username string
	Password string
	Phone    string
	Email    string
}

// @Summary 注册
// @Description 用户注册
// @Param username body string true "Username"
// @Param password body string true "Psername"
// @Param phone body string    false "Phone"
// @Param email body string    false "Email"
// @Accept  json
// @Produce json
// @Success 200 {string} string "hello"
// @Router /user/regis [post]
func (u *UserService) Regis(ctx *gin.Context) {
	ctx.String(200, "abc")
}

// @Summary 用户登录
// @Description 用户登录
// @Produce json
// @Param body body models.UserLoginReq true "body参数"
// @Success 200 {object} models.UserLoginRes "返回用户信息"
// @Router /user/login [post]
func (u *UserService) Login(ctx *gin.Context) {
	ctx.String(200, "abc")
}

// @Summary 用户退出登录
// @Accept  json
// @Produce json
// @Success 200 {string} string "hello"
// @Router /user/logout [post]
func (u *UserService) Logout(ctx *gin.Context) {
	ctx.String(200, "abc")
}

// @Summary 更新用户信息
// @Produce  json
// @parameters  json
// @Success 200 {string} string "hello"
// @Router /user/updateInfo [post]
func (u *UserService) UpdateInfo(ctx *gin.Context) {
	ctx.String(200, "abc")
}

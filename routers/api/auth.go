package api

import (
	"gin-blog/models"
	"gin-blog/pkg/app"
	"gin-blog/pkg/err"
	"gin-blog/pkg/util"
	"gin-blog/service/auth_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	var user models.Auth
	c.ShouldBind(&user)
	username := user.Username
	password := util.Md5(user.Password)

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, err.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{
		Username: username,
		Password: password,
	}
	isExist, e := authService.Check()
	if e != nil {
		appG.Response(http.StatusInternalServerError, err.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, err.ERROR_AUTH, nil)
		return
	}

	token, e := util.GenerateToken(username, password)
	if e != nil {
		appG.Response(http.StatusInternalServerError, err.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, err.SUCCESS, map[string]string{
		"token": token,
	})
}

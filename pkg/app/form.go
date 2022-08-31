package app

import (
	"gin-blog/pkg/err"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	e := c.ShouldBind(form)
	if e != nil {
		return http.StatusBadRequest, err.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, e := valid.Valid(form)
	if e != nil {
		return http.StatusInternalServerError, err.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, err.INVALID_PARAMS
	}

	return http.StatusOK, err.SUCCESS
}

package v1

import (
	"blog-service/global"
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"errors"
	"github.com/gin-gonic/gin"
)

//Validate 参数验证
func Validate(param interface{}, c *gin.Context) (*app.Response, error) {
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return nil, errors.New("invalid param")
	}
	return response, nil
}

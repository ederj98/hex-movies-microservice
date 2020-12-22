package middlewares

import (
	"github.com/fmcarrero/bookstore_utils-go/logger"
	"github.com/fmcarrero/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Validator interface {
	Error() string
	IsBusinessLogic() bool
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		// Use reflect.TypeOf(err.Err) to known the type of your error
		if _, ok := errors.Cause(err.Err).(Validator); ok {
			restErr := rest_errors.NewBadRequestError(err.Error())
			logger.Error(restErr.Message(), restErr)
			c.JSON(restErr.Status(), restErr)
			return
		}
		// Use reflect.TypeOf(err.Err) to known the type of your error
		if _, ok := errors.Cause(err.Err).(exceptions.UserNotFound); ok {
			restErr := rest_errors.NewNotFoundError(err.Error())
			logger.Error(restErr.Message(), restErr)
			c.JSON(restErr.Status(), restErr)
			return
		}
		restErr := rest_errors.NewInternalServerError(err.Error(), err)
		logger.Error(restErr.Message(), restErr)
		c.JSON(restErr.Status(), restErr)
	}
}

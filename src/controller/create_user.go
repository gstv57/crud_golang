package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gstv57/crud_golang/src/configuration/logger"
	"github.com/gstv57/crud_golang/src/configuration/validation/users"
	"github.com/gstv57/crud_golang/src/controller/model/request"
	"github.com/gstv57/crud_golang/src/model"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("error trying to validate user info", err,
			zap.String("Journey", "createUser"))

		errRest := users.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
	}
	c.JSON(http.StatusOK, domain)
}

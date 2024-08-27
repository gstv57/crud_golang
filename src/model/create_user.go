package model

import (
	"github.com/gstv57/crud_golang/src/configuration/logger"
	"github.com/gstv57/crud_golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("Journey", "createUser"))
	ud.EncryptPassword()
	return nil
}

package service

import (
	"fmt"

	"github.com/cedrazgabriel/my-first-crud-go/src/configuration/logger"
	"github.com/cedrazgabriel/my-first-crud-go/src/configuration/rest_err"
	"github.com/cedrazgabriel/my-first-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init create user model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}

package model

import (
	"fmt"

	"github.com/cedrazgabriel/my-first-crud-go/src/configuration/logger"
	"github.com/cedrazgabriel/my-first-crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init create user model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud.Password)

	return nil
}

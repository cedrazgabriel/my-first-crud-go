package model

import "github.com/cedrazgabriel/my-first-crud-go/src/configuration/rest_err"

type userDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

type UserDomainInterface interface {
	CreateUser(userDomain) *rest_err.RestErr
	UpdateUser(string, userDomain) *rest_err.RestErr
	FindUser(string) (*userDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

func (ud *userDomain) CreateUser() *rest_err.RestErr {
	return nil
}

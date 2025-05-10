package model

import "github.com/cedrazgabriel/my-first-crud-go/src/configuration/rest_err"

func (*userDomain) DeleteUser(string) *rest_err.RestErr {
	return nil
}

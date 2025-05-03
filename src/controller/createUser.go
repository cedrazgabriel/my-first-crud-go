package controller

import (
	"github.com/cedrazgabriel/my-first-crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("error creating user")
	c.JSON(err.Code, err)
}

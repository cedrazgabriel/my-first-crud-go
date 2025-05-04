package controller

import (
	"github.com/cedrazgabriel/my-first-crud-go/src/configuration/logger"
	"github.com/cedrazgabriel/my-first-crud-go/src/configuration/validation"
	"github.com/cedrazgabriel/my-first-crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Creating user...",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error binding JSON", err,
			zap.String("journey", "createUser"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

}

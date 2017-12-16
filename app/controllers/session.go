package controllers

import (
	"github.com/gin-gonic/gin"
)

type (
	SessionsController struct{}
	User               struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
)

func NewSessionsController() *SessionsController {
	return &SessionsController{}
}

func (sc SessionsController) CreateSession(c *gin.Context) {
	var user User

	c.BindJSON(&user)

	if user.Username == "username" && user.Password == "password" {
		c.JSON(200, gin.H{"status": "you are logged in"})
	} else {
		c.JSON(401, gin.H{"status": "unauthorized"})
	}
}

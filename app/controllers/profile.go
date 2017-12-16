package controllers

import (
	"github.com/gin-gonic/gin"
)

type (
	ProfilesController struct{}
)

func NewProfilesController() *ProfilesController {
	return &ProfilesController{}
}

func (pc ProfilesController) Show(c *gin.Context) {
	c.JSON(200, gin.H{"status": "User Detail"})
}

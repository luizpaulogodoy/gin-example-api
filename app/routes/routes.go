package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/luizpaulogodoy/gin-example-api/app/controllers"
	"github.com/luizpaulogodoy/gin-example-api/app/middlewares"
)

func Setup(r gin.IRouter) {

	/* ---------------------------  Public routes  --------------------------- */
	public := r.Group("/v1")

	sc := controllers.NewSessionsController()
	public.POST("/login", sc.CreateSession)

	/* ---------------------------  Private routes  --------------------------- */
	private := r.Group("/v1")

	private.Use(middlewares.AuthRequired())
	{
		pc := controllers.NewProfilesController()
		private.GET("/profile", pc.Show)
	}

}

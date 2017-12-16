package main

import (
	"github.com/luizpaulogodoy/gin-example-api/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.Setup(app)

	app.Run(":8080")
}

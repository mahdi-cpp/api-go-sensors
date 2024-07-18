package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddSensorRoutes(rg *gin.RouterGroup) {

	users := rg.Group("/sensor")

	users.GET("/GetTemperature", func(context *gin.Context) {
		var temp = model.Temperature{ID: 12, Temp: 0.5, Name: "Inside  Temperature"}
		context.JSON(200, temp)
	})

	users.GET("/humidity", func(context *gin.Context) {
		context.JSON(http.StatusOK, "humidity sensor...")
	})

	users.GET("/matin", func(context *gin.Context) {
		context.JSON(http.StatusOK, "hello matin...")
	})
}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-sensors/model"
	"net/http"
)

func AddSensorRoutes(rg *gin.RouterGroup) {

	users := rg.Group("/sensor")

	users.GET("/temperature", func(context *gin.Context) {
		var temp = model.Temperature{ID: 12, Temp: 0.5, Name: "Inside  Temperature"}
		context.JSON(200, temp)
	})

	users.GET("/humidity", func(context *gin.Context) {
		context.JSON(http.StatusOK, "humidity sensor...")
	})

	users.GET("/motion-sensor", func(context *gin.Context) {
		context.JSON(http.StatusOK, "hello motion sensor...")
	})
}

package routes

import (
	controller "example.com/alphacoder-server/controllers"
	"example.com/alphacoder-server/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.GET("/sheets/frazsheet", controller.Frazsheet())
	incomingRoutes.GET("/sheets/dsa450sheet", controller.Sheet450())
	incomingRoutes.GET("/sheets/apnasheet", controller.ApnacollegeSheet())
	incomingRoutes.GET("/sheets/blind75sheet", controller.SheetBlind75())
	incomingRoutes.GET("/sheet", controller.AllQuesion())
}

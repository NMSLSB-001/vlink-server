package vlink_server

import (
	"github.com/gin-gonic/gin"
	"vlink-server/internal/app/vlink_server/handler"
)

func SetupRouters(r *gin.Engine) {

	r.POST("/rooms/create", handler.CreateRoom)
	r.GET("/rooms/info/:roomId", handler.GetRoomInfo)
	r.POST("/rooms/join/:roomId", handler.JoinRoom)
	r.POST("/rooms/exit/:roomId", handler.ExitRoom)
	r.POST("/rooms/turn/:roomId", handler.SetTurnInfo)
	r.POST("/rooms/notify/:roomId", handler.SendNotification)
	// r.GET("/healthcheck", handler.HealthCheck)

	// r.GET("", handler.HandleWSNotification)
}

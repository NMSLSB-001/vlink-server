package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vlink-server/internal/app/vlink_server/helper"
	"vlink-server/internal/app/vlink_server/service"
)

func CreateRoom(c *gin.Context) {
	roomId := c.DefaultPostForm("roomId", "")
	if roomId == "" {
		c.JSON(http.StatusOK, helper.FailureJson(helper.RequestEmptyError, "Room ID is required", nil))
		return
	}

	service.CreateRoom(roomId)
	c.JSON(http.StatusOK, helper.SuccessJson(nil))
}

// 获取房间信息
func GetRoomInfo(c *gin.Context) {
	roomId := c.Param("roomId")
	roomInfo := service.GetRoomInfo(roomId)

	if roomInfo == nil {
		c.JSON(http.StatusOK, helper.FailureJson(helper.ErrRoomNotFound, "Room not found", nil))
		return
	}

	c.JSON(http.StatusOK, helper.SuccessJson(roomInfo))
}

func JoinRoom(c *gin.Context) {

}

func ExitRoom(c *gin.Context) {

}

func SetTurnInfo(c *gin.Context) {

}

func SendNotification(c *gin.Context) {

}

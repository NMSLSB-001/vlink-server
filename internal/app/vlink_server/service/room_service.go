package service

import "vlink-server/internal/app/vlink_server/model"

var rooms = make(map[string]*model.Room)

func CreateRoom(roomId string) {

}

func GetRoomInfo(roomId string) *model.Room {

	return rooms[roomId]
}

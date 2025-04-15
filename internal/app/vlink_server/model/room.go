package model

// Room 定义房间信息的结构体
type Room struct {
	RoomId     string `json:"roomId"`     // 房间ID
	RoomSize   int    `json:"roomSize"`   // 房间大小
	Format     string `json:"format"`     // 帧格式设置
	HostPeerId string `json:"hostPeerId"` // 房主的Peer ID
	Turn       string `json:"turn"`       // TURN服务器地址
}

// NewRoom 创建一个新的Room实例
func NewRoom(roomId string, roomSize int, format, hostPeerId, turn string) *Room {
	return &Room{
		RoomId:     roomId,
		RoomSize:   roomSize,
		Format:     format,
		HostPeerId: hostPeerId,
		Turn:       turn,
	}
}

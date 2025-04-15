package model

// RoomInfo 定义房间信息的结构体
type RoomInfo struct {
	RoomId     string `json:"roomId"`     // 房间ID
	RoomSize   int    `json:"roomSize"`   // 房间大小
	Format     string `json:"format"`     // 帧格式设置
	HostPeerId string `json:"hostPeerId"` // 房主的Peer ID
	Turn       string `json:"turn"`       // TURN服务器地址
}

// NewRoomInfo 创建一个新的RoomInfo实例
func NewRoomInfo(roomId string, roomSize int, format, hostPeerId, turn string) *RoomInfo {
	return &RoomInfo{
		RoomId:     roomId,
		RoomSize:   roomSize,
		Format:     format,
		HostPeerId: hostPeerId,
		Turn:       turn,
	}
}

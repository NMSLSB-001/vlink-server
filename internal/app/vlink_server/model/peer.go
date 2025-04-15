package model

// Peer 定义对等方的模型
type Peer struct {
	PeerId       string     `json:"peerId"`       // 对等方的 ID
	IsServer     bool       `json:"isServer"`     // 标识该 Peer 是否是服务器
	Nick         string     `json:"nick"`         // 对等方的昵称
	Rtt          int        `json:"rtt"`          // 往返时间（Round-Trip Time），即网络延迟
	NatType      string     `json:"natType"`      // 网络地址转换类型（NAT 类型）
	SortingOrder int        `json:"sortingOrder"` // 排序顺序，用于排列对等方
	ShareInfo    *ShareInfo `json:"shareInfo"`    // 分享信息（可以为空）
}

// NewPeer 创建一个新的 Peer 实例，包含 ShareInfo
func NewPeer(peerId string, isServer bool, nick string, rtt int, natType string, sortingOrder int, shareInfo *ShareInfo) *Peer {
	return &Peer{
		PeerId:       peerId,
		IsServer:     isServer,
		Nick:         nick,
		Rtt:          rtt,
		NatType:      natType,
		SortingOrder: sortingOrder,
		ShareInfo:    shareInfo,
	}
}

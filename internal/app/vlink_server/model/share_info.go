package model

// ShareInfo 定义分享信息的结构体
type ShareInfo struct {
	Gpu        string `json:"gpu"`        // GPU 信息
	Capture    string `json:"capture"`    // 捕获设备信息
	Start      bool   `json:"start"`      // 是否开始分享
	IsWireless bool   `json:"isWireless"` // 是否无线连接
	Is2G4      bool   `json:"is2G4"`      // 是否 2.4 GHz 网络
}

// NewShareInfo 创建一个新的 ShareInfo 实例
func NewShareInfo(gpu, capture string, start, isWireless, is2G4 bool) *ShareInfo {
	return &ShareInfo{
		Gpu:        gpu,
		Capture:    capture,
		Start:      start,
		IsWireless: isWireless,
		Is2G4:      is2G4,
	}
}

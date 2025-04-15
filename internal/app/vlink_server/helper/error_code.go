package helper

const (
	Success = 0

	RequestEmptyError   = 10001
	ValidateFailedError = 10002

	// 20000 - 20099 RoomInfo
	ErrRoomNotFound = 20001

	// 20100 - 20199 RoomJoin
	ErrJoinUnexpect = 20100
	ErrRoomFull     = 20101

	// 20200 - 20299: RoomCreate
	ErrRoomCreateFail  = 20200
	ErrDuplicateRoomId = 20201

	//30000 - 30099: PeerInfo
	ErrPeerNotFound = 30001
)

package helper

const (
	Success = 0

	RequestEmptyError   = 10001
	ValidateFailedError = 10002
	AuthFailedError     = 10003

	// 30000 - 39999: BasicUser
	ErrUserNotBindExternal    = 30010
	ErrUserUnbindExternal     = 30011
	ErrUserNotAuthorizeScopes = 30012

	// 40000 - 49999: AuthServer
	ErrAppNotFound        = 40001
	ErrInvalidRedirectURI = 40002
	ErrInvalidAppID       = 40003
	ErrInvalidScope       = 40004
	ErrInvalidContext     = 40006
	ErrTokenExpired       = 40007
	ErrInvalidToken       = 40008
	ErrAuthCodeExpired    = 40009
	ErrAuthorizeIDExpired = 40010

	ErrSendPhoneCodeLimit = 40100
	ErrInvalidPhoneCode   = 40101
	ErrPhoneCodeExpired   = 40102

	ErrInvalidProviderType  = 40200
	ErrInvalidState         = 40201
	ErrInvalidAuthorizeCode = 40202

	ErrExternalUserIsBound = 40300
	ErrBasicUserNotFound   = 40301

	ErrAppNotAuthorizeScopes = 40400
)

package helper

type JsonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}

func SuccessJson(data any) *JsonResponse {
	return &JsonResponse{
		Code:    Success,
		Message: "success",
		Data:    data,
	}
}

func FailureJson(code int, msg string, data any) *JsonResponse {
	return &JsonResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

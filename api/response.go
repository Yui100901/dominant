package api

//
// @Author yfy2001
// @Date 2025/1/21 13 57
//

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func NewResponse(code int, data any, message string) *Response {
	return &Response{code, data, message}
}

func Success(data any) *Response {
	return NewResponse(200, data, "OK")
}

func Error(message string) *Response {
	return NewResponse(500, nil, message)
}

package server

//
// @Author yfy2001
// @Date 2024/9/13 16 38
//

// Response 通用的返回结构
type Response struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 消息
	Data interface{} `json:"data"` // 数据
}

// SuccessResponse 创建一个成功的响应
func SuccessResponse(data interface{}) *Response {
	return &Response{
		Code: 200,
		Msg:  "OK",
		Data: data,
	}
}

// ErrorResponse 创建一个错误的响应
func ErrorResponse(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

package handle

type Response struct {
	Code int                    `json:"code"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"msg"`
}

func NewResponse(code int, data map[string]interface{}, msg string) Response {
	return Response{
		Code: code,
		Data: data,
		Msg:  msg,
	}
}

type UserResponse struct {
	UserID     string  `json:"userID"`
	Address    string  `json:"address"`
	Balance    float64 `json:"balance"`
	CreateTime string  `json:"createTime"`
	Invitee    string  `json:"invitee"`
}

type TaskResponse struct {
	UserID  string `json:"userID"`
	Address string `json:"address"`
}

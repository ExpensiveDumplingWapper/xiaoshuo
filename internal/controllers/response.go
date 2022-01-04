package controllers

var (
	SuccessResponse  = Response{Code: 200, Message: "SUCCESS"}
	SuccessNoContent = Response{Code: 204, Message: "SUCCESS"}

	// 全局
	ErrUnauthorized    = Response{Code: 401, Message: "登陆会话失效，请重新登陆"}
	ErrMissParams      = Response{Code: 401, Message: "缺少参数"}
	ErrParams          = Response{Code: 422, Message: "参数校验失败"}
	ErrFailParams      = Response{Code: 400, Message: "参数格式错误"}
	ErrInternalServer  = Response{Code: 500, Message: "系统错误"}
	ErrNotFound        = Response{Code: 404, Message: "未找到"}
	ErrDBServerFailure = Response{Code: 500, Message: "数据库异常，请重试"}
	ErrLoginType       = Response{Code: 403, Message: "登陆类型错误，请尝试更新APP后重试"}
	FailService        = Response{Code: 500, Message: ""}
	ErrNotExist        = Response{Code: 10004, Message: "数据不存在"}
	ErrDefault         = Response{Code: 10005, Message: "操作失败"}
	ErrDataPermission  = Response{Code: 10006, Message: "没有此数据权限"}
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (e Response) Error() string {
	return e.Message
}

func (e Response) SetData(data interface{}) Response {
	e.Data = data
	return e
}

func NewErrResponse(message string) Response {
	s := ErrInternalServer
	s.Message = message
	return s
}

func NewSucResponse(data interface{}) Response {
	s := SuccessResponse
	s.Data = data
	return s
}

func NewErrMissParamsResponse(message string) Response {
	s := ErrMissParams
	s.Message = message
	return s
}

func NewErrFailParamsResponse(message string) Response {
	s := ErrFailParams
	s.Message = message
	s.Data = []string{}
	return s
}

func NewFailResponse(message Response) Response {
	s := message
	s.Data = []string{}
	return s
}

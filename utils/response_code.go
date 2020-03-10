package utils

const (
	RCODE_OK                   = 200 // 请求成功
	RCODE_CEEATED              = 201 // 创建成功
	RCODE_DELETED              = 204 // 删除成功
	RCODE_BAD_REQUEST          = 400 // 请求的地址不存在或者包含不支持的参数
	RECODE_UNAUTHORIZED        = 401 // 未授权
	RECODE_FORBIDDEN           = 403 // 被禁止访问
	RECODE_NOT_FOUND           = 404 // 请求的资源不存在
	RECODE_UNPROCESABLE_ENTITY = 422 // 当创建一个对象时，发生一个验证错误
	RECODE_SERVER_ERROR        = 500 // 服务器内部错误
)

var recodeText = map[int]string{
	RCODE_OK:                   "请求成功",
	RCODE_CEEATED:              "创建成功",
	RCODE_DELETED:              "删除成功",
	RCODE_BAD_REQUEST:          "请求的地址不存在或者包含不支持的参数",
	RECODE_UNAUTHORIZED:        "未授权",
	RECODE_FORBIDDEN:           "被禁止访问",
	RECODE_NOT_FOUND:           "请求的资源不存在",
	RECODE_UNPROCESABLE_ENTITY: "当创建一个对象时，发生一个验证错误",
	RECODE_SERVER_ERROR:        "服务器内部错误",
}

func RecodeText(code int) string {
	if str, ok := recodeText[code]; ok {
		return str
	}
	return ""
}

package error

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_EXIST_TAG:                 "已存在该标签名称",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "token超时",
	ERROR_AUTH_TOKEN:                "token生成失败",
	ERROR_AUTH:                      "token错误",
	ERROR_USER_EXIST:                "该用户名已存在，请修改",
	ERROT_EXIST_USER_NAME:           "该用户名已存在，请修改后提交",
	ERROR_DELETE_CONTENT_NOT_MYSELF: "删除失败，不是本人的评论",
	ERROR_USER_NOT_EXIST:            "该用户不存在",
	ERROR_USER_LOGIN:                "用户名或密码错误",
	ERROR_TODO_DELETE_ERROE:         "todo 删除失败",
	ERROR_TODO_DELETE_BYNAME:        "清空失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

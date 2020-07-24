package error

const (
	SUCCESS                         = 200
	ERROR                           = 500
	INVALID_PARAMS                  = 400
	ERROR_EXIST_TAG                 = 10001
	ERROR_AUTH_CHECK_TOKEN_FAIL     = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT  = 20002
	ERROR_AUTH_TOKEN                = 20003
	ERROR_AUTH                      = 20004
	ERROT_EXIST_USER_NAME           = 402
	ERROR_DELETE_CONTENT_NOT_MYSELF = 408
	ERROR_USER_NOT_EXIST            = 409
	ERROR_USER_LOGIN                = 412
	ERROR_TODO_DELETE_ERROE         = 414
	ERROR_TODO_DELETE_BYNAME        = 415
)

const (
	ERROR_USER_EXIST = 300
)

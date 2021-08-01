package biz_error

type errorType struct {
	Code    int
	Message string
}

type BizError struct {
	errorType
	Error string
}

var (
	MysqlError    = errorType{1000, "读写mysql错误"}
	RedisError    = errorType{1001, "读写redis错误"}
	ResourceError = errorType{1002, "请求资源不存在或无权限"}
	ParamError    = errorType{1003, "参数错误"}
)

func NewMysqlError(err error) *BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &BizError{MysqlError, errMsg}
}

func NewRedisError(err error) *BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &BizError{RedisError, errMsg}
}

func NewResourceError(err error) *BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &BizError{ResourceError, errMsg}
}

func NewParamError(err error) *BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &BizError{ParamError, errMsg}
}

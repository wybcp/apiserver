package errno

import "fmt"

// Errno 错误代码信息
type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

// Err represents an error
type Err struct {
	Code    int
	Message string
	Err     error
}

// New 创建一个新的错误
func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

// Add 增加错误信息
func (err *Err) Add(message string) error {
	//err.Message = fmt.Sprintf("%s %s", err.Message, message)
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	//return err.Message = fmt.Sprintf("%s %s", err.Message, fmt.Sprintf(format, args...))
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// IsErrUserNotFound 没有用户错误
func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

// DecodeErr 解析错误信息
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}

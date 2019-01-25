package errno

var (
	// Common errors

	// OK Common errors
	OK = &Errno{Code: 0, Message: "OK"}
	// InternalServerError 内部错误
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	// ErrBind 添加错误
	ErrBind = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	// ErrValidation Validation failed.
	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	// ErrDatabase Database error.
	ErrDatabase = &Errno{Code: 20002, Message: "Database error."}
	// ErrToken Error occurred while signing the JSON web token.
	ErrToken = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors

	// ErrUserNotFound user errors
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}
	// ErrEncrypt Error occurred while encrypting the user password.
	ErrEncrypt = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	// ErrTokenInvalid The token was invalid.
	ErrTokenInvalid = &Errno{Code: 20103, Message: "The token was invalid."}
	// ErrPasswordIncorrect The password was incorrect.
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
)

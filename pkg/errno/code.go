package errno

const (
	SuccessCode = 0
	ServiceErrCode    = 10001
	ParamParseErrCode = 10002
	ParamErrCode = 10101
	LoginErrCode              = 10202
	UserNotExistErrCode       = 10203
	UserAlreadyExistErrCode   = 10204
	TokenExpiredErrCode       = 10205
	TokenValidationErrCode    = 10206
	TokenInvalidErrCode       = 10207
	UserNameValidationErrCode = 10208
	PasswordValidationErrCode = 10209
	VideoDataCopyErrCode = 10302
	CommentTextErrCode = 10401
	ActionTypeErrCode = 10501
)

var (
	Success = NewErrNo(SuccessCode, "Success")
	ServiceErr    = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamParseErr = NewErrNo(ParamParseErrCode, "Could not parse the param")
	ParamErr = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	LoginErr              = NewErrNo(LoginErrCode, "Wrong username or password")
	UserNotExistErr       = NewErrNo(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr   = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	UserNameValidationErr = NewErrNo(UserNameValidationErrCode, "Username is invalid")
	PasswordValidationErr = NewErrNo(PasswordValidationErrCode, "Password is invalid")
	VideoDataCopyErr = NewErrNo(VideoDataCopyErrCode, "Could not copy video data")
	CommentTextErr = NewErrNo(CommentTextErrCode, "Comment text too long")
	ActionTypeErr = NewErrNo(ActionTypeErrCode, "Action type is invalid")
)

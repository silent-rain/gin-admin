/*业务码消息 简体中文-中国*/
package errcode

import "errors"

// 状态码映射具体消息
var MsgZHCN = map[ErrorCode]error{
	Ok:                   nil,
	InternalError:        errors.New("内部错误"),
	UnknownError:         errors.New("未知错误"),
	HttpServerCloseError: errors.New("http 服务关闭错误"),
	RouteNotFoundError:   errors.New("404 接口不存在"),
	InternalServerError:  errors.New("服务器内部错误"),
	// 定时任务
	TickerRunnerError: errors.New("即时器执行错误"),
	TickerPanicError:  errors.New("即时器严重错误"),
	TimerRunnerError:  errors.New("定时器执行错误"),
	TimerPanicError:   errors.New("定时器严重错误"),
	// 请求
	ReqParameterParsingError:    errors.New("请求参数解析错误"),
	ReqContentTypeNotFoundError: errors.New("请求 Content-Type 参数不存在"),
	ReqContentTypeParamsError:   errors.New("请求 Content-Type 参数错误"),
	// 数据解析
	JsonDataEncodeError: errors.New("数据编码错误"),
	JsonDataDecodeError: errors.New("数据解码错误"),
	// 数据库
	DBQueryError:             errors.New("数据查询错误"),
	DBQueryEmptyError:        errors.New("数据不存在"),
	DBAddError:               errors.New("数据添加失败"),
	DBUpdateError:            errors.New("数据更新失败"),
	DBDeleteError:            errors.New("数据删除失败"),
	DBBatchDeleteError:       errors.New("数据批量删除失败"),
	DBUpdateStatusError:      errors.New("更新状态失败"),
	DBResetError:             errors.New("数据重置失败"),
	DBDataExistError:         errors.New("数据已存在"),
	DBDataExistChildrenError: errors.New("存在子项"),
	DBWriteCloseError:        errors.New("读写数据库实例关闭失败"),
	DBReadCloseError:         errors.New("只读数据库实例关闭失败"),
	// 鉴权
	TokenGenerateError: errors.New("生成 Token 失败"),
	TokenNotFound:      errors.New("鉴权信息不存在"),
	TokenParsingError:  errors.New("解析 Token 失败"),
	TokeConvertError:   errors.New("转换 Token 失败"),
	TokenInvalidError:  errors.New("无效鉴权"),
	TokenExpiredError:  errors.New("鉴权过期"),
	// 系统管理
	UserRegisterError:           errors.New("用户注册失败"),
	UserLoginError:              errors.New("用户登录失败"),
	UserLogoutError:             errors.New("用户注销失败"),
	UserDisableError:            errors.New("您的账号已被禁用,请联系管理员"),
	UserNamePasswordError:       errors.New("账号或密码不正确"),
	UserOldPasswordError:        errors.New("旧密码不正确"),
	UserPasswordError:           errors.New("密码不正确"),
	UserPhoneConsistentError:    errors.New("新旧手机号码一致, 未更新"),
	UserEmailConsistentError:    errors.New("新旧邮箱一致, 未更新"),
	CaptchaTypeError:            errors.New("验证码类型错误, 不支持的类型"),
	CaptchaEtxNotFoundError:     errors.New("验证码格式异常"),
	CaptchaNotFoundError:        errors.New("验证码不存在"),
	CaptchaGenerateError:        errors.New("生成验证码失败"),
	CaptchaVerifyError:          errors.New("验证码错误"),
	SessionGetCaptchaEmptyError: errors.New("验证码为空"),
	ExistPhoneError:             errors.New("手机号已存在"),
	ExistEmailError:             errors.New("邮箱已存在"),
	// 文件上传
	UploadFileParserError: errors.New("上传文件解析失败"),
	UploadFileSaveError:   errors.New("上传文件保存失败"),
	// 系统操作
	FileNotFoundError: errors.New("文件不存在"),
	DirNotFoundError:  errors.New("文件夹不存在"),
}

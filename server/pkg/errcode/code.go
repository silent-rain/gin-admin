// Package errcode 业务状态码
package errcode

import "fmt"

// 业务状态码
type ErrorCode uint

const (
	Ok                   ErrorCode = iota + 10000 // 访问成功
	InternalError                                 // 内部错误
	UnknownError                                  // 未知错误
	HttpServerCloseError                          // Http 服务关闭错误
	RouteNotFoundError                            // 404 接口不存在
	InternalServerError                           // 500 服务器内部错误
	LoadSSLError                                  // 加载 SSL 错误
	LoadIndexHtmlError                            // 加载首页文件失败
)

// 定时任务
const (
	TickerRunnerError ErrorCode = iota + 10050 // 即时器执行错误
	TickerPanicError                           // 即时器严重错误
	TimerRunnerError                           // 定时器执行错误
	TimerPanicError                            // 定时器严重错误
)

// 请求
const (
	ReqParameterParsingError    ErrorCode = iota + 10100 // 请求参数解析错误
	ReqContentTypeNotFoundError                          // 请求 Content-Type 参数不存在
	ReqContentTypeParamsError                            // 请求 Content-Type 参数错误
)

// 数据解析
const (
	JsonDataEncodeError     ErrorCode = iota + 10200 // json 数据编码错误
	JsonDataDecodeError                              // json 数据解码错误
	DataTypeConversionError                          // 数据类型转换失败
)

// Mysql 数据库
const (
	DBQueryError             ErrorCode = iota + 10300 // 数据库查询错误
	DBQueryEmptyError                                 // 数据不存在
	DBAddError                                        // 数据添加失败
	DBUpdateError                                     // 数据更新失败
	DBUpsertError                                     // 数据插入或更新失败
	DBDeleteError                                     // 数据删除失败
	DBBatchDeleteError                                // 数据批量删除失败
	DBUpdateStatusError                               // 更新状态失败
	DBResetError                                      // 数据重置失败
	DBDataExistError                                  // 数据已存在
	DBDataExistChildrenError                          // 存在子项
	DBWriteCloseError                                 // 读写数据库实例关闭失败
	DBReadCloseError                                  // 只读数据库实例关闭失败
)

// Redis 数据库
const (
	RedisPingError           ErrorCode = iota + 10350 // redis 连接失败
	RedisUnknownClientError                           // redis 未知客户端
	RedisSetKeyError                                  // redis 设置 KEY 失败
	RedisGetKeyError                                  // redis 获取 KEY 失败
	RedisDelKeyError                                  // redis 删除 KEY 失败
	RedisIncrKeyError                                 // redis 自增 KEY 值失败
	RedisDecrKeyError                                 // redis 自减 KEY 值失败
	RedisTTLGetKeyError                               // redis 查看 Key 剩余的过期时间失败
	RedisSetKeyExpireError                            // redis 设置 key 的过期时间失败
	RedisSetKeyExpireAtError                          // redis 设置 key 的过期时间戳失败
	RedisGetKeyExistsError                            // redis 查看一个或多个 KEY 是否存在异常
)

// Cache 内部缓存
const (
	CacheGetError ErrorCode = iota + 10380 // 获取缓存失败
)

// 鉴权
const (
	TokenGenerateError                ErrorCode = iota + 10400 // 生成 Token 失败
	TokenNotFound                                              // 鉴权信息不存在
	TokenParsingError                                          // 解析 Token 失败
	TokeConvertError                                           // 转换 Token 失败
	TokenInvalidError                                          // 无效鉴权
	TokenExpiredError                                          // 鉴权过期
	TokenIssuerError                                           // 签发人错误
	TokenUnconformityError                                     // Token 不一致, 已在别处登录，请重新登录
	TokenDisableCurrentLoginError                              // 当前登录已被禁用
	ApiHttpTokenPassphraseEmptyError                           // 令牌口令不能为空
	ApiHttpTokenPassphraseError                                // 令牌口令错误
	ApiHttpTokenInvalidError                                   // 无效API令牌
	ApiHttpTokenMethodPermissionError                          // 没有该资源的操作权限
)

// 上游服务

// 系统管理
const (
	UserRegisterError           ErrorCode = iota + 11000 // 用户注册失败
	UserLoginError                                       // 用户登录失败
	UserLogoutError                                      // 用户注销失败
	UserDisableError                                     // 您的账号已被禁用,请联系管理员
	UserNamePasswordError                                // 账号或密码不正确
	UserOldPasswordError                                 // 旧密码不正确
	UserPasswordError                                    // 密码不正确
	UserPhoneConsistentError                             // 新旧手机号码一致, 未更新
	UserEmailConsistentError                             // 新旧邮箱一致, 未更新
	CaptchaTypeError                                     // 验证码类型错误, 不支持的类型
	CaptchaEtxNotFoundError                              // 验证码格式异常
	CaptchaNotFoundError                                 // 验证码不存在
	CaptchaGenerateError                                 // 生成验证码失败
	CaptchaVerifyError                                   // 验证码错误
	SessionGetCaptchaEmptyError                          // 验证码为空
	ExistPhoneError                                      // 手机号已存在
	ExistEmailError                                      // 邮箱已存在
)

// 文件上传
const (
	UploadFileParserError ErrorCode = iota + 11100 // 上传文件解析失败
	UploadFileSaveError                            // 上传文件保存失败
)

// 系统操作
const (
	FileNotFoundError ErrorCode = iota + 11200 // 文件不存在
	DirNotFoundError                           // 文件夹不存在
	DirCreateError                             // 文件夹创建失败
)

// CodeError 返回状态码错误信息
func (e ErrorCode) CodeError() error {
	msg, ok := MsgZHCN[e]
	if !ok {
		return MsgZHCN[UnknownError]
	}
	return msg
}

// Error 实现 error 接口
func (e ErrorCode) Error() string {
	if e == Ok {
		return ""
	}
	return e.CodeError().Error()
}

// WithMsg 添加额外错误信息
func (e ErrorCode) WithMsg(msg string) *ErrorMsg {
	return &ErrorMsg{
		Code: e,
		Err:  fmt.Errorf("%w, %s", e.CodeError(), msg),
	}
}

// WithMsg 添加额外错误信息
func (e ErrorCode) WithError(err error) *ErrorMsg {
	return &ErrorMsg{
		Code: e,
		Err:  fmt.Errorf("%w, %s", e.CodeError(), err),
	}
}

// Error 自定义错误类型
type ErrorMsg struct {
	Code ErrorCode `json:"code"` // 状态码
	Err  error     `json:"msg"`  // 状态码信息
}

// Error 实现 error 接口
func (e *ErrorMsg) Error() string {
	return e.Err.Error()
}

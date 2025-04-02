package core

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type XContext[T any] struct {
	echo.Context
	gorm     *gorm.DB
	validate *Validator
}

func NewSkipGormGlobalHookContext() context.Context {
	ctx = context.Background()
	return context.WithValue(ctx, GormGlobalSkipHookKey, true)
}

// GetContext  第一个泛型是入参的类型
func GetContext[T any](c echo.Context) *XContext[T] {
	cc := XContext[T]{
		c, GetGormDB(), GetValidator(),
	}
	return &cc
}

// GetHeardParam 获取请求头参数
func (c *XContext[V]) GetHeardParam(key string) string {
	return c.Context.Request().Header.Get(key)
}

// GetAppPlatformCode  获取请求头参数
func (c *XContext[V]) GetAppPlatformCode() string {
	return c.GetHeardParam(AppPlatformHeaderKey)
}

// GetLoginUser  获取请求头参数
func (c *XContext[V]) GetLoginUser() ClaimsAdditions {
	claims, err := GetTokenManager().ParseJwt(c.GetUserToken())
	if err != nil {
		c.CheckError(NewErrCodeMsg(TOKEN_EXPIRE_ERROR, "登录身份过期，请重新登录！"))
		c.CheckError(err)
	}
	return claims.ClaimsAdditions
}

// GetUserToken  获取请求头参数
func (c *XContext[V]) GetUserToken() string {
	param := c.GetHeardParam(Authorization)
	split := strings.Split(param, " ")
	if len(split) != 2 {
		c.CheckError(NewErrCode(TOKEN_FORMAT_ERROR))
	}
	return split[1]
}

// GetLoginUerName GetHeardParam 获取请求头参数
func (c *XContext[V]) GetLoginUerName() string {
	return c.GetLoginUser().NickName
}

// GetLoginUserUid  获取UID
func (c *XContext[V]) GetLoginUserUid() int64 {
	return c.GetLoginUser().UID
}

// GetLoginUserDepartmentId  获取用户所属的部门ID
func (c *XContext[V]) GetLoginUserDepartmentId() int64 {
	return c.GetLoginUser().DepartmentId
}

func (c *XContext[V]) GetQueryParamAndValid() (body V) {
	t := new(V)
	err := c.Bind(t)
	err = c.validate.ValidateStruct(t)
	c.CheckError(err)
	return *t
}
func (c *XContext[V]) GetBodyAndValid() (body V) {
	// 创建一个新的 V 类型的实例
	t := new(V)
	// 绑定请求体到 t
	err := c.Bind(t)
	if err != nil {
		c.CheckError(err)
		return
	}
	// 使用反射检查 t 是否是切片或数组
	val := reflect.ValueOf(t).Elem() // 获取指针指向的实际值
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		// 如果是切片或数组，遍历每个元素并验证
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i).Interface() // 获取当前元素
			err = c.validate.ValidateStruct(elem)
			if err != nil {
				c.CheckError(err)
				return
			}
		}
	} else {
		// 如果不是切片或数组，直接验证 t
		err = c.validate.ValidateStruct(t)
		if err != nil {
			c.CheckError(err)
			return
		}
	}
	// 返回绑定并验证后的值
	return *t
}

// SetEchoContext 中间件函数，设置自定义上下文
func SetEchoContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 创建自定义上下文
		cc := XContext[any]{
			c, GetGormDB(), GetValidator(),
		}
		// 调用下一个处理程序，传递自定义上下文
		return next(cc)
	}
}

func (c *XContext[V]) GetDB() *gorm.DB {
	return c.gorm
}

func (c *XContext[V]) Success(body interface{}) error {
	if !c.Response().Committed {
		return c.JSON(http.StatusOK, c.CreateSuccess(body))
	}
	return nil
}

func (c *XContext[V]) Fail(err *CodeError) error {
	if !c.Response().Committed {
		return c.JSON(http.StatusOK, c.CreateError(err.GetErrCode(), err.GetErrMsg()))
	}
	return nil
}
func (c *XContext[V]) ValidateStruct(s interface{}) error {
	return c.validate.ValidateStruct(s)
}

func (c *XContext[V]) CheckError(err error) {
	if err != nil {
		c.Context.Error(err)
		panic(err)
	}
}

func (c *XContext[V]) QueryArray(queryName string) (arr []string) {
	param := c.QueryParam(queryName)
	return strings.Split(param, ",")
}

// QueryInt64Array parses a query parameter and returns an array of int64
func (c *XContext[V]) QueryInt64Array(queryName string) (arr []int64) {
	split := c.QueryArray(queryName)
	for _, str := range split {
		if str == "" {
			continue // Skip empty values
		}
		value, _ := strconv.ParseInt(str, 10, 64)
		arr = append(arr, value)
	}

	return arr
}
func (c *XContext[V]) GetPathParam(pathName string) string {
	return c.Param(pathName)
}
func (c *XContext[V]) GetPathParamInt64(pathName string) int64 {
	param := c.Param(pathName)
	parseInt, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		c.CheckError(NewFrontShowErrMsg("请输入正确的参数！"))
	}
	return parseInt
}

// QueryParamIds 获取query参数的id列表 参数名为 ids
func (c *XContext[V]) QueryParamIds() []int64 {
	t := new(QueryIds)
	err := c.Bind(t)
	err = c.validate.ValidateStruct(t)
	c.CheckError(err)
	return t.Ids
}

func (c *XContext[V]) CreateSuccess(v any) *ResponseSuccess {
	if v == nil {
		return &ResponseSuccess{
			Code: 200,
			Msg:  "SUCCESS",
		}
	}
	return &ResponseSuccess{
		Code: 200,
		Msg:  "SUCCESS",
		Data: v,
	}
}

func (c *XContext[V]) CreateError(code uint32, msg string) *ResponseError {
	return &ResponseError{
		Code: code,
		Msg:  msg,
	}
}

type ResponseSuccess struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
type ResponseError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func (c *XContext[V]) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

func (c *XContext[V]) Done() <-chan struct{} {
	return nil
}

func (c *XContext[V]) Err() error {
	return nil
}

func (c *XContext[V]) Value(key any) any {
	return nil
}

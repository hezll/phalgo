package phalgo

//	PhalGo-engine
//	注意路由引擎,依赖Echo对器进行封装
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//			"github.com/labstack/echo"

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/engine/standard"
	"net/http"
)

var Echo *echo.Echo

// 初始化echo实例
func NewEcho() *echo.Echo {

	Echo = echo.New()
	return Echo
}

// 使用Fasthttp方式开启服务
func RunFasthttp(prot string) {

	Echo.Run(fasthttp.New(prot))
}

// 使用Standard的方式开启服务
func RunStandard(prot string) {

	Echo.Run(standard.New(prot))
}

// 打印请求异常信息
func Recover() {

	Echo.Use(middleware.Recover())
}

// 打印请求信息
func Logger() {

	Echo.Use(middleware.Logger())
}

// 开启gzip压缩
func Gzip() {

	Echo.Use(middleware.Gzip())
}

// 自动添加末尾斜杠
func AddTrailingSlash() {

	Echo.Use(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
}

// 自动删除末尾斜杠
func RemoveTrailingSlash() {

	Echo.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
}

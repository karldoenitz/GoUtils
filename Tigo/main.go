package main

import (
    "github.com/karldoenitz/Tigo/TigoWeb"
    "net/http"
)

// DemoHandler handler
type DemoHandler struct {
    TigoWeb.BaseHandler
}

func (demoHandler *DemoHandler) Get() {
    demoHandler.ResponseAsText("Hello Demo!")
}

// Authorize 中间件
func Authorize(w *http.ResponseWriter, r *http.Request) bool {
	// 此处返回true表示继续执行，false则直接返回，后续的中间件不会执行
	return true
}

// 路由
var urls = []TigoWeb.Pattern{
    {"/demo", DemoHandler{}, []TigoWeb.Middleware{Authorize}},
}

func main() {
    application := TigoWeb.Application{
        IPAddress:   "127.0.0.1",
        Port:        8888,
        UrlPatterns: urls,
    }
    application.Run()
}

package router

import (
	"fmt"
	"net/http"
)

type Router struct {

}

// 用router实现http.Handler必须要实现ServeHTTP方法
func (r *Router) ServeHTTP(w http.ResponseWriter,req *http.Request) {
	switch req.URL.Path {
	case "/Good":
		fmt.Fprint(w,"is good")
	case "/Nice":
		fmt.Fprint(w,"is Nice")
	case "/ANODVNIUENGINDS":
		fmt.Fprint(w,"is luangxie")
	default:
		http.Error(w,"404 Not Found",404)
	}
}
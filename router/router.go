package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {

}

// 用最基本的router实现http.Handler必须要实现ServeHTTP方法
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

type MRouter struct {
	r *mux.Router
}

func GetRouter() MRouter {
	var router MRouter
	router.r = mux.NewRouter()
	return router
}

//func (m *MRouter)
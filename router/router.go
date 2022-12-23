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

func Hello(w http.ResponseWriter,r *http.Request){
	username := r.Context().Value("username").(string)
	fmt.Fprintf(w,"Hi %s\n",username)
}

func GetRouter() *mux.Router {
	var router MRouter
	router.r = mux.NewRouter()
	// http://127.0.0.1:8000/users/asd
	router.r.HandleFunc("/users/{user:[a-z]+}", func(w http.ResponseWriter, request *http.Request) {
		user := mux.Vars(request)["user"]
		fmt.Fprint(w,user)
	}).Methods("GET")

	// http://127.0.0.1:8000/test/才可以访问成功，而http://127.0.0.1:8000/test会失败
	// 因为github.com/gorilla/mux完全正则
	router.r.HandleFunc("/test/", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w,mux.Vars(request))
	}).Methods("GET")

	// http://127.0.0.1:8000/demo
	router.r.HandleFunc("/demo", Hello).Methods("GET")
	return router.r
}

//func (m *MRouter)

var rrouter *mux.Router
func GetMuxRouter() *mux.Router {
	if rrouter == nil {
		rrouter = mux.NewRouter()
		rrouter.HandleFunc("/demo", Hello).Methods("GET")
	}

	return rrouter
}
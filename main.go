package main

import (
	"fmt"
	"mhttpsvr/router"
	"net/http"
)

// fmt.Fprintf 根据格式对io write写入
func HelloMsg(w http.ResponseWriter ,r *http.Request)  {
	// get http://127.0.0.1:8000/hello?name=me
	fmt.Fprintf(w,"Http %s\n",r.URL.Query().Get("name"))
}

func main() {
	fmt.Println("Hello World")
	//var r router.Router
	//var l middleware.Logger
	//l.InitLogger(&r)
	//
	//err := http.ListenAndServe(":8000",&l)
	//if err != nil {
	//	fmt.Println(err)
	//}

	http.ListenAndServe(":8000",router.GetRouter())
}
package mhttpsvr

import (
	"fmt"
	"net/http"
)

// fmt.Fprintf 根据格式对io write写入
func HelloMsg(w http.ResponseWriter ,r *http.Request)  {
	fmt.Fprintf(w,"Http %s\n",r.URL.Query().Get("name"))
}

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/hello",HelloMsg)
	http.ListenAndServe("19000",nil)
}
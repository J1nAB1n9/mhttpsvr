package middleware

import (
	"github.com/urfave/negroni"
	"log"
	"mhttpsvr/router"
	"net/http"
)

type Logger struct {
	Inner http.Handler
}

func (l *Logger)InitLogger(i *router.Router) {
	l.Inner = i
}

func (l *Logger)ServeHTTP(w http.ResponseWriter,r *http.Request) {
	log.Println("start")
	l.Inner.ServeHTTP(w,r)
	log.Println("finish")
}

var Middleware,SimpleMiddleware *negroni.Negroni

func InitMiddleware() {
	// negroni.Classic() 会创建几个默认的中间件
	Middleware = negroni.Classic()
	// negroni.New() 不会创建任何的中间件，除非你往里面传参
	SimpleMiddleware = negroni.New()

	Middleware.UseHandler(router.GetMuxRouter())
}
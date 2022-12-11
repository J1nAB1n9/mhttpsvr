package middleware

import (
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
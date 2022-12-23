package middleware

import (
	"context"
	"net/http"
)

type Auth struct {
	UserName string
	PassWord string
}

func (a *Auth) ServeHTTP(w http.ResponseWriter,r *http.Request,next http.HandlerFunc) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	if a.UserName != username || a.PassWord != password {
		http.Error(w,"username or password err",401)
		return
	}

	ctx := context.WithValue(r.Context(),"username",username)
	r = r.WithContext(ctx)
	next(w,r)
}

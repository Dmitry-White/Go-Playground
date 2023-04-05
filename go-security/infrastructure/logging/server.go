package main

import (
	"expvar"
	"fmt"
	"go-security/infrastructure/logging/services/auth"
	"net/http"
)

var (
	loginCalls = expvar.NewInt("login.calls")
	goodLogins = expvar.NewInt("login.good")
	badLogins  = expvar.NewInt("login.bad")
)

func loginHandler(resw http.ResponseWriter, req *http.Request) {
	appLogger.Info("Help")

	if req.Method != http.MethodPost {
		appLogger.Error("bad method")
		http.Error(resw, "bad method", http.StatusMethodNotAllowed)
		return
	}

	login := req.FormValue(auth.FormUser)
	password := req.FormValue(auth.FormPassword)
	if len(login) == 0 || len(password) == 0 {
		appLogger.Error("missing auth")
		http.Error(resw, "missing auth", http.StatusUnauthorized)
		return
	}

	loginCalls.Add(1)

	user, ok := auth.LoginUser(login, password)
	if !ok {
		badLogins.Add(1)
		appLogger.Error(fmt.Sprintf("bad %q login from %s", login, req.RemoteAddr))
		http.Error(resw, "bad auth", http.StatusUnauthorized)
		return
	}
	goodLogins.Add(1)

	appLogger.Info(fmt.Sprintf("%q login from %s", login, req.RemoteAddr))
	auth.SetUserCookie(resw, user)

	fmt.Fprintf(resw, "Welcome %s!\n", user.Name)
}

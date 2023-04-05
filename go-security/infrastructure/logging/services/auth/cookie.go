package auth

import (
	"net/http"
	"time"
)

func SetUserCookie(resw http.ResponseWriter, u User) {
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:    u.Name,
		Value:   u.Login,
		Expires: expiration,
	}

	http.SetCookie(resw, &cookie)
}

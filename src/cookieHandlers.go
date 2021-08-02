package main

import "net/http"

func setJwtCookie(token string, w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name: token,
	}
	http.SetCookie(w, &c)
}

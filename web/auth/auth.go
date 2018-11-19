package auth

import (
	"log"
	"net/http"
	"strings"
	"time"

	"../../web/auth/storage/memory"
)

// jwtCustomClaims are custom claims extending default ones.
func DoAuthSignup(r *http.Request, w http.ResponseWriter) (string, string) {
	auth := r.Header.Get("Authorization")
	log.Print("Auth Header ->")
	log.Print(r.Header)
	str := strings.Split(auth, ":")
	str2 := strings.Split(str[1], " ")
	un, pw := str2[0], str2[1]

	for _, j := range mymem.Users {
		if strings.Compare(un, j.Username) == 0 {
			return "", ""
		}
	}

	log.Print("Added User Cookie")
	http.SetCookie(w, &http.Cookie{
		Name:    "userInfo",
		Value:   un + ":" + pw,
		Expires: time.Now().Add(1 * time.Hour),
	})

	return un, pw
}

func DoAuthLogin(w http.ResponseWriter, r *http.Request) bool {
	auth := r.Header.Get("Authorization")
	str := strings.Split(auth, ":")
	str2 := strings.Split(str[1], " ")
	un, pw := str2[0], str2[1]
	mymem.SetCurrentUser(un, pw)
	log.Print(auth)
	stat := checkUser(un, pw)
	if stat == true {
		log.Print("VALID")
		http.SetCookie(w, &http.Cookie{
			Name:    "userInfo",
			Value:   un + ":" + pw,
			Expires: time.Now().Add(1 * time.Hour),
		})
		//w.WriteHeader(302)
		return true
	} else {
		//w.WriteHeader(202)
		return false
	}
}
func checkUser(un string, pw string) bool {
	log.Print("Entering checkUser")
	log.Print(un, " ", pw, "\n")
	log.Print(len(mymem.Users))

	//m := make(map[string]string)
	for _, j := range mymem.Users {
		if strings.Compare(un, j.Username) == 0 {
			if strings.Compare(pw, j.Password) == 0 {
				return true
			}
		}
	}

	return false
}

package main

import (
	//"fmt"

	"fmt"
	"html/template"
	"log"
	"net/http"
	//"os"
	//"strings"
)

type post struct {
	Author string
	Desc   string
}

var postsList []post

func initPost() {
	postsList = []post{
		{"Kavya", "We built this city"},
		{"Nikhila", "Favorite Radio City"},
		{"Navi", "On Rock and Roll"},
	}

}

type user_info struct {
	username string
	password string
}

var users []user_info

func basic_user_details() {
	users := []user_info{
		{"nikhila", "hello"},
		{"ubeee", "yippee"},
		{"poorna", "noiceee"},
	}

	fmt.Print(users)
}

func checkUser(un string, pw string) bool {
	var status bool
	for _, j := range users {
		u := j.username
		fmt.Printf(u)
		p := j.password
		fmt.Printf(p)
		if un == u && pw == p {
			status = true
		}
	}
	return status
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	/* if r.Method == "POST" {
		r.ParseForm()
		newPost := post{r.FormValue("author"), r.FormValue("desc")}
		postsList = append(postsList, newPost)
		log.Print("postList in POST: ", postsList)
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	} */
	//log.Print("postList in GET: ", postsList)
	t := template.Must(template.New("login").ParseFiles("../views/login.html"))
	if r.Method == "POST" {
		r.ParseForm()
		loggedUser := user_info{r.FormValue("uname"), r.FormValue("password")}
		status := checkUser(string(loggedUser.username), string(loggedUser.password))

		log.Print(status)
		if status == true {
			log.Print("VALID")
		}
		log.Print("Uname and pwd", loggedUser)
		//ts := template.Must(template.New("posts").ParseFiles("../views/posts.html"))
		http.Redirect(w, r, "/posts.html", http.StatusSeeOther)
	}
	err := t.ExecuteTemplate(w, "login.html", postsList)
	if err != nil {
		log.Fatal("Some error: ", err)
	}

}

func main() {
	initPost()
	basic_user_details()
	log.Print("Calling init")
	log.Print("Calling basic user details")
	http.HandleFunc("/", PostHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

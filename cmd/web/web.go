package main

import (
	//"fmt"
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

func PostHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseForm()
		newPost := post{r.FormValue("author"), r.FormValue("desc")}
		postsList = append(postsList, newPost)
		log.Print("postList in POST: ", postsList)
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	}
	log.Print("postList in GET: ", postsList)
	t := template.Must(template.New("posts").ParseFiles("../views/post.html"))
	err := t.ExecuteTemplate(w, "post.html", postsList)
	if err != nil {
		log.Fatal("Some error: ", err)
	}

}

func main() {
	initPost()
	log.Print("Calling init")
	http.HandleFunc("/", PostHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

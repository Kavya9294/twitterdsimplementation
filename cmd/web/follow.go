package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
	//"os"
	//"strings"
)

type user struct {
	username  string
	userno    int
	followers []int
}

var userList []user

func initPost() {
	userList = []user{
		{"User1", 1, []int{2, 3}},
		{"User2", 2, []int{1}},
		{"User3", 3, []int{2}},
	}
}

func followHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseForm()
		suser := r.FormValue("suser")
		duser := r.FormValue("duser")
		suserVal := 0
		duserVal := 0
		for i := 0; i<len(userList); i++ {
			if(userList[i].username == suser)
			{
				suser := i
			}
			if(userList[i].username == duser)
			{
				duserVal := userList[i].userno
			}
		}
		userList[suser].followers := append(userList[suser].followers, duserVal)
		postsList = append(postsList, newPost)
		log.Print("userList in POST: ", userList)
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	}
	log.Print("userList in GET: ", userList)
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

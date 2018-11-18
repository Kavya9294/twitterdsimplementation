package main

import (
	//"fmt"

	"html/template"
	"log"
	"net/http"

	"../../web/auth"
	"../../web/auth/storage/memory"
	"github.com/gorilla/mux"
	//"os"
	//"strings"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Printf("Entering signup %d", len(mymem.Users))
		un, pw := auth.DoAuthSignup(r)
		mymem.AddUser(un, pw)
		log.Print("New users     -> ", mymem.Users)

	}
	w.WriteHeader(200)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.New("login").ParseFiles("../views/login.html"))
	if r.Method == "POST" {
		r.ParseForm()
		log.Printf("Entering Login Handler POST %d", len(mymem.Users))
		log.Print("Header ->")
		log.Print(r.Header)
		auth.DoAuthLogin(w, r)

	}

	err := t.ExecuteTemplate(w, "login.html", mymem.Users)
	if err != nil {
		log.Fatal("Some error: ", err)
	}
	w.WriteHeader(200)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	all_posts := mymem.PostsList
	log.Print("all posts: ", all_posts)
	users := mymem.Users
	cur_user := users[0]
	cur_posts := cur_user.GetPosts(all_posts)

	if r.Method == "POST" {
		r.ParseForm()
		newPost := mymem.Post{r.FormValue("author"), r.FormValue("desc")}
		all_posts = newPost.AppendPost()
		log.Print("postList in POST: ", all_posts)
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	}
	log.Print("Posts for current user: ", cur_posts)
	t := template.Must(template.New("posts").ParseFiles("../views/post.html"))
	err := t.ExecuteTemplate(w, "post.html", cur_posts)
	if err != nil {
		log.Fatal("Some error: ", err)
	}

}

func main() {
	mymem.Initialize()
	log.Print("Calling init")

	r := mux.NewRouter()
	r.HandleFunc("/signup", SignupHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/posts", PostHandler)
	http.Handle("/", r)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

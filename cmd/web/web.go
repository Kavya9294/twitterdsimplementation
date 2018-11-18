package main

import (
	//"fmt"

	"html/template"
	"log"
	"net/http"

	"../../web/auth"
	"../../web/auth/storage/memory"
	//"os"
	//"strings"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Printf("Entering signup %d", len(mymem.Users))
		un, pw := auth.DoAuthSignup(r, w)
		if un == "" {
			log.Print("Duplicate username ")
			w.WriteHeader(202)
		} else {
			mymem.AddUser(un, pw)
			log.Print("New users     -> ", mymem.Users)
		}
	}
	//w.WriteHeader(200)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.New("login").ParseFiles("../views/login.html"))

	if r.Method == "POST" {
		r.ParseForm()
		log.Printf("Entering Login Handler POST %d", len(mymem.Users))
		log.Print("Header ->")
		log.Print(r.Header)
		ok := auth.DoAuthLogin(w, r)
		//r.Method = "GET"
		if ok {
			http.Redirect(w, r, "/posts", http.StatusFound)
		}
	}
	if r.Method == "GET" {
		err := t.ExecuteTemplate(w, "login.html", mymem.Users)
		if err != nil {
			log.Fatal("Some error: ", err)
		}
	}

}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	cur_user := mymem.GetCurrentUser(r)

	if cur_user.Username == "" {
		log.Fatal("User not authorized")
		// Redirect to login
	} else {
		all_posts := mymem.PostsList
		log.Print("all posts: ", all_posts)
		log.Print("current_user: ", cur_user)
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
}

func main() {
	mymem.Initialize()
	log.Print("Calling init")

	http.HandleFunc("/signup", SignupHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/posts", PostHandler)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package main

import (
	"../../web/auth/storage/memory"
	"html/template"
	"log"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {

	all_posts := *mymem.PostsList
	log.Print("all posts: ", all_posts)
	users := *mymem.Users
	cur_user := users[0]
	cur_posts := cur_user.GetPosts(&all_posts)

	if r.Method == "POST" {
		r.ParseForm()
		newPost := &mymem.Post{r.FormValue("author"), r.FormValue("desc"), cur_user.UserId}
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
	http.HandleFunc("/", PostHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

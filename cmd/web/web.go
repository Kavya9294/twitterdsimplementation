package main

import (
	//"fmt"

	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
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
		{"nikhila", "Favorite Radio City"},
		{"Navi", "On Rock and Roll"},
	}

}

type user_info struct {
	username string
	password string
}

var users []user_info

func basic_user_details() {
	users = []user_info{
		{"nikhila", "hello"},
		{"ubeee", "yippee"},
		{"poorna", "noiceee"},
	}

	fmt.Print(users)
	for _, j := range users {
		fmt.Print(j.username + " ")
		fmt.Print(j.password)
		fmt.Println("")
	}
	log.Printf(" hii %d", len(users))
}

func checkUser(un string, pw string) bool {
	log.Print("Entering checkUser")
	log.Print(un, " ", pw, "\n")
	log.Print(len(users))

	//m := make(map[string]string)
	for _, j := range users {
		if strings.Compare(un, j.username) == 0 {
			if strings.Compare(pw, j.password) == 0 {
				return true
			}
		}
	}

	return false
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
		//basic_user_details()
		if strings.Compare(r.FormValue("LOGIN"), "") >= 0 {
			log.Printf("hello %d", len(users))
			loggedUser := user_info{r.FormValue("uname"), r.FormValue("password")}
			stat := checkUser(loggedUser.username, loggedUser.password)

			log.Print(stat)
			if stat == true {
				log.Print("VALID")
				http.Redirect(w, r, "/posts.html", http.StatusSeeOther)
				r.ParseForm()
				newPost := post{r.FormValue("author"), r.FormValue("desc")}
				postsList = append(postsList, newPost)
				log.Print("postList in POST: ", postsList)
			}
			log.Print("Uname and pwd", loggedUser.username, loggedUser.password)
		}
		if strings.Compare(r.FormValue("SIGNUP"), "") >= 0 {
			log.Printf("hello %d", len(users))
			newUser := user_info{r.FormValue("susername"), r.FormValue("spassword")}
			users = append(users, newUser)
			log.Print("users     -> ", users)
		}

		//ts := template.Must(template.New("posts").ParseFiles("../views/posts.html"))
		//http.Redirect(w, r, "/posts.html", http.StatusSeeOther)
	}
	err := t.ExecuteTemplate(w, "login.html", postsList)
	if err != nil {
		log.Fatal("Some error: ", err)
	}

}

func main() {
	initPost()
	basic_user_details()
	log.Print(len(users))
	log.Print("Calling init")
	log.Print("Calling basic user details")
	http.HandleFunc("/", PostHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

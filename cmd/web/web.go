package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strings"

	"../../web/auth"
	pb "../../web/auth/authpb"
	"google.golang.org/grpc"
)

var client pb.AccessClient
var listOfAllUsers *pb.Users
var listofAllPosts *pb.Posts
var curUserPosts *pb.Posts

func addUser(name string, password string, followers []string) *pb.Users {
	usr := &pb.User{
		Username:  name,
		Password:  password,
		Followers: followers,
	}
	uList, addErr := client.AddUser(context.Background(), usr)

	if addErr != nil {
		log.Fatalf("unable to add user: %v", addErr)
	}
	for _, u := range uList.UsersList {
		log.Println("username: %v, Password: %v", u.Username, u.Password)
	}
	return uList
}

func addPost(name string, desc string) *pb.Posts {
	post := &pb.Post{
		Username: name,
		Desc:     desc,
	}
	pList, addErr := client.AddPost(context.Background(), post)

	if addErr != nil {
		log.Fatalf("unable to add post: %v", addErr)
	}
	for _, p := range pList.PostsList {
		log.Println("username: %v, Desc: %v", p.Username, p.Desc)
	}
	return pList
}

func getCurrentUserPosts(username string) *pb.Posts {
	usr := &pb.User{
		Username: username,
	}
	pList, addErr := client.GetPosts(context.Background(), usr)

	if addErr != nil {
		log.Fatalf("unable to get Post: %v", addErr)
	}
	log.Println("Current User Posts")
	for _, p := range pList.PostsList {
		log.Println("username: %v, Desc: %v", p.Username, p.Desc)
	}
	return pList

}

func getAllUsers(username string) *pb.Users {
	usr := &pb.User{
		Username: username,
	}
	uList, getErr := client.GetAllUsers(context.Background(), usr)

	if getErr != nil {
		log.Fatalf("unable to get all users due to: %v", getErr)
	}
	log.Println("All users")
	for _, u := range uList.UsersList {
		log.Println("username: %v", u.Username)
	}
	return uList

}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var t *template.Template

	//HardCode
	useName := "Kavya"
	allUsers := getAllUsers(useName)

	t = template.Must(template.New("login").ParseFiles("../views/login.html"))

	if r.Method == "GET" {
		err := t.ExecuteTemplate(w, "login.html", allUsers)
		if err != nil {
			log.Fatal("Some error: ", err)
		}
	}
	if r.Method == "POST" {
		//log.Printf("Entering signup %d", len(mymem.Users))
		un, pw := auth.DoAuthSignup(r, w, allUsers)

		if un == "" {
			log.Print("Duplicate username ")
			w.WriteHeader(401)
		} else {
			addUser(un, pw, []string{un})
			w.WriteHeader(302)
			cur_user, err := client.SetCurrentUser(context.Background(), &pb.User{
				Username:  un,
				Password:  pw,
				Followers: []string{un},
			})
			if err != nil {
				log.Println("err,cannot assign current user due to: ", err)
			}
			log.Print("Current user     -> ", cur_user.CurUser.Username)
		}
	}

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var t *template.Template

	t = template.Must(template.New("login").ParseFiles("../views/login.html"))

	if r.Method == "GET" {
		AllUsersForLogin := getAllUsers("login")
		err := t.ExecuteTemplate(w, "login.html", AllUsersForLogin)
		if err != nil {
			log.Fatal("Some error: ", err)
		}
	}

	if r.Method == "POST" {
		r.ParseForm()
		//log.Printf("Entering Login Handler POST %d", len(mymem.Users))
		log.Print("Header ->")
		log.Print(r.Header)
		allLoginUsers := getAllUsers("LOgin")
		ok, un, pw := auth.DoAuthLogin(w, r, allLoginUsers)
		if ok {
			//log.Print("current_user: ", mymem.Cur_user)
			cur_user, err := client.SetCurrentUser(context.Background(), &pb.User{
				Username:  un,
				Password:  pw,
				Followers: []string{un},
			})
			if err != nil {
				log.Println("error in login")
			}
			log.Println(cur_user)
			w.WriteHeader(302)
			return
		} else {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
		}
	}

}

func getFollowers(username string) []string {
	usr := &pb.User{
		Username: username,
	}
	uList, getErr := client.GetAllUsers(context.Background(), usr)

	var userList []string
	if getErr != nil {
		log.Println("could not get all users due to error: ", getErr)
	} else {

		for _, user := range uList.UsersList {
			if strings.Compare(user.Username, username) != 0 {
				userList = append(userList, user.Username)
			}
		}
	}
	return userList

}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	//Need to split auth header each time and get current user
	cur_user, err := client.GetCurrentUser(context.Background(), &pb.User{Username: "kavya"})

	log.Print("current_user: ", cur_user.CurUser.Username)

	if err != nil {
		log.Printf("User not authorized")
		// Redirect to login
	} else {
		//all_posts := mymem.PostsList
		//log.Print("all posts: ", all_posts)
		Following := getFollowers(cur_user.CurUser.Username)
		log.Print("Following: ", Following)
		cur_posts := getCurrentUserPosts(cur_user.CurUser.Username)

		log.Println("Current Users Posts: ", cur_posts)

		if r.Method == "POST" {
			r.ParseForm()
			//newPost := &pb.Post{
			//Username: cur_user.CurUser.Username,
			//Desc:     r.FormValue("desc"),
			//}
			all_posts := addPost(cur_user.CurUser.Username, r.FormValue("desc"))
			log.Print("postList in POST: ", all_posts)
			http.Redirect(w, r, "/post", http.StatusSeeOther)
		}
		log.Print("Posts for current user: ", cur_posts)
		paths := []string{
			"../views/post.html",
			"../views/following.html",
		}
		var t *template.Template
		t = template.Must(template.ParseFiles(paths...))
		log.Print("t: ", t)

		type Response struct {
			Following []string
			Posts     *pb.Posts
		}
		var response Response
		response.Following = Following
		response.Posts = cur_posts
		log.Println("response.Posts: ", response.Posts.PostsList)
		err := t.ExecuteTemplate(w, "post.html", response)
		if err != nil {
			log.Fatal("Some error: ", err)
		}
	}
}

//func FollowsHandler(w http.ResponseWriter, r *http.Request) {
//log.Print("in Follows Handler")
//Following := mymem.Cur_user.Following
//log.Print("following: ", Following)
//toggle_user := strings.TrimPrefix(r.URL.Path, "/follows/")
//followers_list := mymem.ToggleFollower(toggle_user)
//log.Print("followers_list: ", followers_list)
//http.Redirect(w, r, "/post", http.StatusSeeOther)

//}

func initialize() {

	addUser("Nikhila", "aGVsbG8=", []string{"Nikhila", "Kavya", "Nikhila"})
	addUser("Kavya", "eWlwcGVl", []string{"Kavya", "Navi"})
	listOfAllUsers = addUser("Navi", "bm9pY2VlZQ==", []string{"Navi", "Nikhila"})

	addPost("Nikhila", "Life is great")
	addPost("Kavya", "Music is Life")
	addPost("Navi", "Rock and roll all the way")
	addPost("Navi", "Pink floyed-Wish you were here-#rythm#to#ears")
	addPost("Nikhila", "Artic Monkeys#best#ever#music")
	listofAllPosts = addPost("Kavya", "Traveller mode ON #One#Life")
	curUserPosts = getCurrentUserPosts("Kavya")
}

func main() {
	log.Print("Calling init")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	client = pb.NewAccessClient(conn)
	initialize()

	http.HandleFunc("/", SignupHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/post", PostHandler)
	//http.HandleFunc("/follows/", FollowsHandler)

	server_err := http.ListenAndServe(":9090", nil)
	if server_err != nil {
		log.Fatal("ListenAndServe: ", server_err)
	}
}

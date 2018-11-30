package main

import (
	pb "../../web/auth/authpb"
	//"../../web/auth"
	"context"
	"google.golang.org/grpc"
	//"html/template"
	"log"
	//"net/http"
	//"strings"
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

//func SignupHandler(w http.ResponseWriter, r *http.Request) {
//if r.Method == "POST" {
//log.Printf("Entering signup %d", len(mymem.Users))
//un, pw := auth.DoAuthSignup(r, w)
//if un == "" {
//log.Print("Duplicate username ")
//w.WriteHeader(401)
//} else {
//mymem.AddUser(un, pw)
//w.WriteHeader(302)

//log.Print("New users     -> ", mymem.Users)
//}
//}

//}

//func LoginHandler(w http.ResponseWriter, r *http.Request) {

//var t *template.Template

//t = template.Must(template.New("login").ParseFiles("../views/login.html"))

//if r.Method == "GET" {
//err := t.ExecuteTemplate(w, "login.html", mymem.Users)
//if err != nil {
//log.Fatal("Some error: ", err)
//}
//}

//if r.Method == "POST" {
//r.ParseForm()
//log.Printf("Entering Login Handler POST %d", len(mymem.Users))
//log.Print("Header ->")
//log.Print(r.Header)
//ok := auth.DoAuthLogin(w, r)
//if ok {
//log.Print("current_user: ", mymem.Cur_user)
//w.WriteHeader(302)
//return
//} else {
//http.Redirect(w, r, "/login", http.StatusUnauthorized)
//}
//}

//}

//func PostHandler(w http.ResponseWriter, r *http.Request) {

//Following := mymem.GetAllUsers()

//if mymem.Cur_user.Username == "" {
//log.Printf("User not authorized")
//// Redirect to login
//} else {
//all_posts := mymem.PostsList
//log.Print("all posts: ", all_posts)
//log.Print("current_user: ", mymem.Cur_user)
//cur_posts := mymem.Cur_user.GetPosts(all_posts)

//if r.Method == "POST" {
//r.ParseForm()
//newPost := mymem.Post{mymem.Cur_user.Username, r.FormValue("desc")}
//all_posts = mymem.AppendPost(newPost)
//log.Print("postList in POST: ", all_posts)
//http.Redirect(w, r, "/post", http.StatusSeeOther)
//}
//log.Print("Posts for current user: ", cur_posts)
//paths := []string{
//"../views/post.html",
//"../views/following.html",
//}
//var t *template.Template
//t = template.Must(template.ParseFiles(paths...))
//log.Print("t: ", t)

//type Response struct {
//Following []string
//Posts     []mymem.Post
//}
//var response Response
//response.Following = Following
//response.Posts = cur_posts
//err := t.ExecuteTemplate(w, "post.html", response)
//if err != nil {
//log.Fatal("Some error: ", err)
//}
//}
//}

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

	addUser("Nikhila", "1234", []string{"Nikhila", "Kavya", "Nikhila"})
	addUser("Kavya", "1234", []string{"Kavya", "Navi"})
	listOfAllUsers = addUser("Navi", "1234", []string{"Navi", "Nikhila"})

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

	//http.HandleFunc("/signup", SignupHandler)
	//http.HandleFunc("/login", LoginHandler)
	//http.HandleFunc("/post", PostHandler)
	//http.HandleFunc("/follows/", FollowsHandler)

	//err := http.ListenAndServe(":9090", nil)
	//if err != nil {
	//log.Fatal("ListenAndServe: ", err)
	//}
}

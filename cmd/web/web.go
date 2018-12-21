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

	t = template.Must(template.New("login").ParseFiles("../views/login.html"))

	if r.Method == "GET" {

		err := t.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			log.Fatal("Some error: ", err)
		}
	}
	if r.Method == "POST" {
		log.Println("In POST of SIGNUP")
		//requester := getReqesterName(r)
		//log.Println("Requester: ", requester)
		allUsers := getAllUsers("CurUser")

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
	log.Println("In GET of POST")
	requester := getReqesterName(r)
	log.Println("Requester: ", requester)
	cur_user, err := client.GetCurrentUser(context.Background(), &pb.User{Username: requester})

	log.Print("current_user: ", cur_user.CurUser.Username)

	if err != nil {
		log.Printf("User not authorized")
		// Redirect to login
	} else {
		Following := getFollowers(cur_user.CurUser.Username)
		log.Print("Following: ", Following)
		cur_posts := getCurrentUserPosts(cur_user.CurUser.Username)

		log.Println("Current Users Posts: ", cur_posts)

		if r.Method == "POST" {
			r.ParseForm()
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

func FollowsHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("in Follows Handler")
	//Following := mymem.Cur_user.Following
	//log.Print("following: ", Following)
	toggleUser := strings.TrimPrefix(r.URL.Path, "/follows/")
	destUser := &pb.User{
		Username: toggleUser,
	}
	log.Println("dest User : ", toggleUser)
	requester := getReqesterName(r)
	log.Println("requester: ", requester)
	cur_user, cerr := client.GetCurrentUser(context.Background(), &pb.User{Username: requester})

	if cerr != nil {
		log.Printf("Error")
	}

	log.Print("current_user: ", cur_user.CurUser.Username)

	FUser := &pb.FollowUser{
		SourceUser: cur_user,
		DestUser:   destUser,
	}
	_, e := client.ToggleFollowers(context.Background(), FUser)
	if e != nil {
		log.Printf("Error")
	}
	http.Redirect(w, r, "/post", http.StatusSeeOther)
}

func getReqesterName(r *http.Request) string {
	cookie, err := r.Cookie("userInfo")

	if err != nil {
		log.Print("Error in getREquesterName : ", err)
	}
	userInfo := cookie.Value
	temp_string := strings.Split(userInfo, ":")
	un := temp_string[0]
	return un
}

func main() {
	log.Print("Calling init")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	client = pb.NewAccessClient(conn)
	client.Initialise(context.Background(), &pb.User{
		Username: "init",
	})

	http.HandleFunc("/", SignupHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/post", PostHandler)
	http.HandleFunc("/follows/", FollowsHandler)

	server_err := http.ListenAndServe(":9090", nil)
	if server_err != nil {
		log.Fatal("ListenAndServe: ", server_err)
	}
}

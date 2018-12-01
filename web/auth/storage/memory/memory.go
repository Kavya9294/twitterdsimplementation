package main

import "fmt"

import (
	"context"
	"log"
	"net"

	pb "../../authpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	listOfPosts []*pb.Post
	listOfUsers []*pb.User
	currentUser *pb.User
}

func (s *server) AddUser(ctx context.Context, in *pb.User) (*pb.Users, error) {

	s.listOfUsers = append(s.listOfUsers, in)

	resp := new(pb.Users)
	resp.UsersList = s.listOfUsers
	return resp, nil
}

func (s *server) GetPosts(ctx context.Context, in *pb.User) (*pb.Posts, error) {
	fmt.Print("In get Posts")
	user := new(pb.User)
	currentUserName := in.Username
	for _, i := range s.listOfUsers {
		if i.Username == currentUserName {
			user = i
		}
	}
	followingPosts := new(pb.Posts)
	all_posts := s.listOfPosts
	for _, following := range user.Followers {
		for _, ind_post := range all_posts {
			if ind_post.Username == following {

				followingPosts.PostsList = append(followingPosts.PostsList, ind_post)
				fmt.Print("Posts in followingPosts: ", followingPosts)
			}
		}
	}
	return followingPosts, nil
}

func (s *server) AddPost(ctx context.Context, in *pb.Post) (*pb.Posts, error) {
	s.listOfPosts = append(s.listOfPosts, in)
	resp := new(pb.Posts)
	resp.PostsList = s.listOfPosts
	return resp, nil
}

func (s *server) SetCurrentUser(ctx context.Context, in *pb.User) (*pb.CurrentUser, error) {
	user := new(pb.CurrentUser)
	user.CurUser = in
	s.currentUser = in
	return user, nil
}

func (s *server) GetCurrentUser(ctx context.Context, in *pb.User) (*pb.CurrentUser, error) {
	user := new(pb.CurrentUser)
	//for _, usr := range s.listOfUsers {
	//if in.Username == usr.Username {
	//user.CurUser = usr
	//}
	//}
	user.CurUser = s.currentUser
	return user, nil
}

//func GetCurrentUser(req *http.Request) User {
//cookie, err := req.Cookie("userInfo")
//fmt.Print("Cookie: ", cookie)
//if err != nil {
//fmt.Print("Error in getCurrentUser : ", err)
//return User{}
//}
//userInfo := cookie.Value
//fmt.Print("userInfo: ", userInfo)
//temp_string := strings.Split(userInfo, ":")
//un, pw := temp_string[0], temp_string[1]
//for _, user := range Users {
//if un == user.Username && pw == user.Password {
//Cur_user = user
//return Cur_user
//}
//}
//return User{"", "", []string{""}}

//}

//func ToggleFollower(duser string) []string {
//pos := -1
//following_list := Cur_user.Following
//var following_new_list []string
//for index, following := range following_list {
//if duser == following {
//pos = index
//}
//}
//if pos == -1 {
//following_list = append(following_list, duser)
//following_new_list = following_list
//} else {
//for i, follow := range following_list {
//if i != pos {
//following_new_list = append(following_new_list, follow)
//}
//}
//}
//Cur_user.Following = following_new_list
//return following_new_list
//}

func (s *server) GetAllUsers(ctx context.Context, in *pb.User) (*pb.Users, error) {
	user_list := new(pb.Users)
	//userName := in.CurUser.Username
	//for _, user := range s.listOfUsers {
	//if userName != user.Username {
	//user_list.UsersList = append(user_list.UsersList, user)
	//}
	//}
	user_list.UsersList = s.listOfUsers
	return user_list, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccessServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

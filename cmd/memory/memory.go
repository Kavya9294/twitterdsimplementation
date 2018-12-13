package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"go.etcd.io/etcd/clientv3"

	pb "../../web/auth/authpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//"go.etcd.io/etcd/raft/raftpb"

const (
	port = ":50051"
)

type server struct {
	listOfPosts []*pb.Post
	listOfUsers []*pb.User
	currentUser *pb.User
}
type lUser struct {
	Username  string
	Password  string
	Followers []string
}
type uList struct {
	UsersList []lUser
}
type lPost struct {
	Username string
	Desc     string
}

type uPosts struct {
	PostsList []lPost
}

type currentU stuct {
	CurUser lUser
}

type followU struct {
	SourceUser currentU
	DestUser lUser
}

func (s *server) Initialise(ctx context.Context, in *pb.User) (*pb.Users, error) {
	var flag1, flag2 bool

	if s.listOfUsers == nil {
		s.listOfUsers = []*pb.User{
			&pb.User{Username: "Nikhila", Password: "aGVsbG8=", Followers: []string{"Nikhila", "Kavya"}},
			&pb.User{Username: "Kavya", Password: "eWlwcGVl", Followers: []string{"Kavya", "Navi"}},
			&pb.User{Username: "Navi", Password: "bm9pY2VlZQ==", Followers: []string{"Navi", "Nikhila"}},
		}
		flag1 = true
	} else {
		flag1 = false
	}
	if s.listOfPosts == nil {
		s.listOfPosts = []*pb.Post{
			&pb.Post{Username: "Nikhila", Desc: "Life is great"},
			&pb.Post{Username: "Kavya", Desc: "Music is Life"},
			&pb.Post{Username: "Navi", Desc: "Rock and roll all the way"},
			&pb.Post{Username: "Navi", Desc: "Pink floyed-Wish you were here-#rythm#to#ears"},
			&pb.Post{Username: "Nikhila", Desc: "Artic Monkeys#best#ever#music"},
		}
		flag2 = true
	} else {
		flag2 = false
	}
	user_resp := new(pb.Users)
	if flag1 && flag2 {
		log.Println("Initialisation complete")
		log.Println(s.listOfUsers)
		log.Println(s.listOfPosts)

		user_resp.UsersList = s.listOfUsers
		return user_resp, nil
	}
	return user_resp, nil
}

func (s *server) AddUser(ctx context.Context, in *pb.User) (*pb.Users, error) {

	log.Print("In addUser")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	resp, rerr := cli.Get(ctx, "User")

	var u_List uList

	if rerr != nil {
		log.Print("Error: ", rerr)
	} else {
		for _, ev := range resp.Kvs {
			fmt.Printf("Value:  %s\n ", ev.Value)
			_ = json.Unmarshal(ev.Value, &u_List)
			fmt.Print("Json thing\n", u_List.UsersList)
		}

	}

	log.Print("usrList: ", u_List.UsersList)
	tempUser := lUser{
		Username:  in.Username,
		Password:  in.Password,
		Followers: in.Followers,
	}
	log.Print("Temp user: ", tempUser)
	u_List.UsersList = append(u_List.UsersList, tempUser)
	log.Print("u_list: ", u_List.UsersList)
	marred, _ := json.Marshal(u_List)
	log.Print("marred: ", string(marred))
	_, err = cli.Put(ctx, "User", string(marred))

	cancel()
	u_List2 := new(pb.Users)

	for _, u := range u_List.UsersList {
		t := &pb.User{
			Username:  u.Username,
			Password:  u.Password,
			Followers: u.Followers,
		}
		print("t.Followes: ", t.Followers)
		u_List2.UsersList = append(u_List2.UsersList, t)
	}

	return u_List2, nil
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
	//s.currentUser = in
	return user, nil
}

func (s *server) GetCurrentUser(ctx context.Context, in *pb.User) (*pb.CurrentUser, error) {
	user := new(pb.CurrentUser)
	for _, usr := range s.listOfUsers {
		if in.Username == usr.Username {
			user.CurUser = usr
		}
	}

	//user.CurUser = s.currentUser
	return user, nil
}

func (s *server) ToggleFollowers(ctx context.Context, in *pb.FollowUser) (*pb.User, error) {
	user := new(pb.User)
	for _, i := range s.listOfUsers {
		if i.Username == in.SourceUser.CurUser.Username {
			user = i
		}
	}
	pos := -1
	following_list := user.Followers
	var following_new_list []string
	for index, following := range following_list {
		if in.DestUser.Username == following {
			pos = index
		}
	}
	if pos == -1 {
		following_list = append(following_list, in.DestUser.Username)
		following_new_list = following_list
	} else {
		for i, follow := range following_list {
			if i != pos {
				following_new_list = append(following_new_list, follow)
			}
		}
	}
	in.SourceUser.CurUser.Followers = following_new_list
	for i := 0; i < len(s.listOfUsers); i++ {
		if s.listOfUsers[i].Username == in.SourceUser.CurUser.Username {
			s.listOfUsers[i] = in.SourceUser.CurUser
			break
		}
	}

	return in.SourceUser.CurUser, nil
}

func (s *server) GetAllUsers(ctx context.Context, in *pb.User) (*pb.Users, error) {
	user_list := new(pb.Users)
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

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

type currentU struct {
	CurUser lUser
}

type followU struct {
	SourceUser currentU
	DestUser   lUser
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

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	resp, rerr := cli.Get(ctx, "Post")

	var all_posts uPosts
	if rerr != nil {
		log.Print("Error in posts: ", rerr)
	} else {
		for _, ev := range resp.Kvs {
			fmt.Printf("Value:  %s\n ", ev.Value)
			_ = json.Unmarshal(ev.Value, &all_posts)
			fmt.Print("Json thing Post\n", all_posts)
		}
	}

	var temp_user lUser
	currentUserName := in.Username
	respUser, respErr := cli.Get(ctx, "User")
	var u_List uList

	if respErr != nil {
		log.Print("Error: ", rerr)
	} else {
		for _, ev := range respUser.Kvs {
			fmt.Printf("Value:  %s\n ", ev.Value)
			_ = json.Unmarshal(ev.Value, &u_List)
			fmt.Print("Json thing user\n", u_List.UsersList)
		}
	}

	for _, i := range u_List.UsersList {
		if i.Username == currentUserName {
			temp_user = i
			break
		}
	}
	followingPosts := new(pb.Posts)

	for _, following := range temp_user.Followers {
		for _, ind_post := range all_posts.PostsList {
			if ind_post.Username == following {
				temp_post := &pb.Post{
					Username: ind_post.Username,
					Desc:     ind_post.Desc,
				}
				followingPosts.PostsList = append(followingPosts.PostsList, temp_post)
				fmt.Print("Posts in followingPosts: ", followingPosts)
			}
		}
	}
	cancel()
	return followingPosts, nil
}

func (s *server) AddPost(ctx context.Context, in *pb.Post) (*pb.Posts, error) {
	//s.listOfPosts = append(s.listOfPosts, in)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	resp, rerr := cli.Get(ctx, "Post")

	var all_posts uPosts
	new_post := lPost{
		Username: in.Username,
		Desc:     in.Desc,
	}
	if rerr != nil {
		log.Print("Error in posts: ", rerr)
	} else {
		for _, ev := range resp.Kvs {
			fmt.Printf("Value:  %s\n ", ev.Value)
			_ = json.Unmarshal(ev.Value, &all_posts)
			fmt.Print("Json thing Post\n", all_posts)
		}
	}
	all_posts.PostsList = append(all_posts.PostsList, new_post)
	marred, _ := json.Marshal(all_posts)
	_, err = cli.Put(ctx, "Post", string(marred))

	cancel()
	all_posts2 := new(pb.Posts)
	for _, p := range all_posts.PostsList {
		t := &pb.Post{
			Username: p.Username,
			Desc:     p.Desc,
		}
		//print("t.Followes: ", t.Followers)
		all_posts2.PostsList = append(all_posts2.PostsList, t)
	}

	return all_posts2, nil

	//r := new(pb.Posts)
	//r.PostsList = s.listOfPosts
	//return r, nil
}

func (s *server) SetCurrentUser(ctx context.Context, in *pb.User) (*pb.CurrentUser, error) {
	user := new(pb.CurrentUser)
	user.CurUser = in
	return user, nil
}

func (s *server) GetCurrentUser(ctx context.Context, in *pb.User) (*pb.CurrentUser, error) {
	log.Print("\nIn getCureentUser\n")

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
			fmt.Print("Json thing CurrentUser\n", u_List.UsersList)
		}

	}
	user := new(pb.CurrentUser)

	for _, usr := range u_List.UsersList {
		if in.Username == usr.Username {
			t_user := &pb.User{
				Username:  usr.Username,
				Password:  usr.Password,
				Followers: usr.Followers,
			}
			user.CurUser = t_user
			break
		}
	}
	cancel()
	return user, nil
}

func (s *server) ToggleFollowers(ctx context.Context, in *pb.FollowUser) (*pb.User, error) {
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
			fmt.Print("Json thing Follows\n", u_List.UsersList)
		}

	}

	user := new(pb.User)
	for _, i := range u_List.UsersList {
		if i.Username == in.SourceUser.CurUser.Username {
			user.Username = i.Username
			user.Password = i.Password
			user.Followers = i.Followers
			break
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
	var temp_list pb.Users
	var u_List2 uList

	for _, usr := range u_List.UsersList {
		u := &pb.User{
			Username:  usr.Username,
			Password:  usr.Password,
			Followers: usr.Followers,
		}
		log.Print("USers each time: ", u)
		if usr.Username == in.SourceUser.CurUser.Username {
			tUsr := in.SourceUser.CurUser
			temp_list.UsersList = append(temp_list.UsersList, tUsr)
			new_usr := lUser{
				Username:  tUsr.Username,
				Password:  tUsr.Password,
				Followers: tUsr.Followers,
			}
			u_List2.UsersList = append(u_List2.UsersList, new_usr)
		} else {
			temp_list.UsersList = append(temp_list.UsersList, u)
			u_List2.UsersList = append(u_List2.UsersList, usr)
		}
	}

	marred, _ := json.Marshal(u_List2)
	log.Print("marred: ", string(marred))
	_, err = cli.Put(ctx, "User", string(marred))

	cancel()

	return in.SourceUser.CurUser, nil
}

func (s *server) GetAllUsers(ctx context.Context, in *pb.User) (*pb.Users, error) {
	//user_list := new(pb.Users)
	//user_list.UsersList = s.listOfUsers
	//return user_list, nil

	log.Print("In GetAllUsers")
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

	log.Print("all users in get all users: ", u_List2)

	return u_List2, nil

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

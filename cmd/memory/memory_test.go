package main

import (
	"context"
	"reflect"
	"testing"

	pb "../../web/auth/authpb"
)

/*
func Test_server_Initialise(t *testing.T) {
	go func() {
		lis, _ := net.Listen("tcp", port)
		s := grpc.NewServer()
		pb.RegisterAccessServer(s, &server{})
		reflection.Register(s)
		s.Serve(lis)
	}()

	type fields struct {
		listOfPosts []*pb.Post
		listOfUsers []*pb.User
		currentUser *pb.User
	}
	type args struct {
		ctx context.Context
		in  *pb.User
	}
	testUserList := []*pb.User{
		&pb.User{Username: "Nikhila", Password: "aGVsbG8=", Followers: []string{"Nikhila", "Kavya"}},
		&pb.User{Username: "Kavya", Password: "eWlwcGVl", Followers: []string{"Kavya", "Navi"}},
		&pb.User{Username: "Navi", Password: "bm9pY2VlZQ==", Followers: []string{"Navi", "Nikhila"}},
	}
	testPostList := []*pb.Post{
		&pb.Post{Username: "Nikhila", Desc: "Life is great"},
		&pb.Post{Username: "Kavya", Desc: "Music is Life"},
		&pb.Post{Username: "Navi", Desc: "Rock and roll all the way"},
		&pb.Post{Username: "Navi", Desc: "Pink floyed-Wish you were here-#rythm#to#ears"},
		&pb.Post{Username: "Nikhila", Desc: "Artic Monkeys#best#ever#music"},
	}
	testCurrentUser := &pb.User{
		Username:  "Nikhila",
		Password:  "aGVsbG8=",
		Followers: []string{"Nikhila", "Kavya"},
	}
	testpbUser := &pb.User{
		Username: "Nikhila",
	}
	testwant := new(pb.Users)
	testwant.UsersList = testUserList
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Users
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Initialize Test1", fields{testPostList, testUserList, testCurrentUser}, args{context.Background(), testpbUser}, testwant, false},
	}
	go func() {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := &server{
					listOfPosts: tt.fields.listOfPosts,
					listOfUsers: tt.fields.listOfUsers,
					currentUser: tt.fields.currentUser,
				}
				got, err := s.Initialise(tt.args.ctx, tt.args.in)
				if (err != nil) != tt.wantErr {
					t.Errorf("server.Initialise() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("server.Initialise() = %v, want %v", got, tt.want)
				}
			})
		}
	}()
}
*/
func Test_server_AddUser(t *testing.T) {
	type fields struct {
		listOfPosts []*pb.Post
		listOfUsers []*pb.User
		currentUser *pb.User
	}
	type args struct {
		ctx context.Context
		in  *pb.User
	}
	testUserList := []*pb.User{
		&pb.User{Username: "Nikhila", Password: "aGVsbG8=", Followers: []string{"Nikhila", "Kavya"}},
		&pb.User{Username: "Kavya", Password: "eWlwcGVl", Followers: []string{"Kavya", "Navi"}},
		&pb.User{Username: "Navi", Password: "bm9pY2VlZQ==", Followers: []string{"Navi", "Nikhila"}},
	}

	testPostList := []*pb.Post{
		&pb.Post{Username: "Nikhila", Desc: "Life is great"},
		&pb.Post{Username: "Kavya", Desc: "Music is Life"},
		&pb.Post{Username: "Navi", Desc: "Rock and roll all the way"},
		&pb.Post{Username: "Navi", Desc: "Pink floyed-Wish you were here-#rythm#to#ears"},
		&pb.Post{Username: "Nikhila", Desc: "Artic Monkeys#best#ever#music"},
	}
	newUserList := []*pb.User{
		&pb.User{Username: "Nikhila", Password: "aGVsbG8=", Followers: []string{"Nikhila", "Kavya"}},
		&pb.User{Username: "Kavya", Password: "eWlwcGVl", Followers: []string{"Kavya", "Navi"}},
		&pb.User{Username: "Navi", Password: "bm9pY2VlZQ==", Followers: []string{"Navi", "Nikhila"}},
		&pb.User{Username: "Nikhila2", Password: "aGVsbG8=3", Followers: []string{"Kavya", "Navi"}},
	}
	testCurrentUser := &pb.User{
		Username:  "Nikhila",
		Password:  "aGVsbG8=",
		Followers: []string{"Nikhila", "Kavya"},
	}
	newUser := &pb.User{
		Username:  "Nikhila2",
		Password:  "aGVsbG8=3",
		Followers: []string{"Kavya", "Navi"},
	}
	testwant := new(pb.Users)
	testwant.UsersList = newUserList
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Users
		wantErr bool
	}{
		// TODO: Add test cases.
		{"AddUserTest1", fields{testPostList, testUserList, testCurrentUser}, args{context.Background(), newUser}, testwant, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				listOfPosts: tt.fields.listOfPosts,
				listOfUsers: tt.fields.listOfUsers,
				currentUser: tt.fields.currentUser,
			}
			got, err := s.AddUser(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_GetPosts(t *testing.T) {
	type fields struct {
		listOfPosts []*pb.Post
		listOfUsers []*pb.User
		currentUser *pb.User
	}
	type args struct {
		ctx context.Context
		in  *pb.User
	}
	testUserList := []*pb.User{
		&pb.User{Username: "Nikhila", Password: "aGVsbG8=", Followers: []string{"Nikhila", "Kavya"}},
		&pb.User{Username: "Kavya", Password: "eWlwcGVl", Followers: []string{"Kavya", "Navi"}},
		&pb.User{Username: "Navi", Password: "bm9pY2VlZQ==", Followers: []string{"Navi", "Nikhila"}},
	}

	testPostList := []*pb.Post{
		&pb.Post{Username: "Nikhila", Desc: "Life is great"},
		&pb.Post{Username: "Kavya", Desc: "Music is Life"},
		&pb.Post{Username: "Navi", Desc: "Rock and roll all the way"},
		&pb.Post{Username: "Navi", Desc: "Pink floyed-Wish you were here-#rythm#to#ears"},
		&pb.Post{Username: "Nikhila", Desc: "Artic Monkeys#best#ever#music"},
	}
	testCurrentUser := &pb.User{
		Username:  "Nikhila",
		Password:  "aGVsbG8=",
		Followers: []string{"Nikhila", "Kavya"},
	}
	postsUser := &pb.User{
		Username: "Nikhila",
	}
	returnpostsList := []*pb.Post{
		&pb.Post{Username: "Nikhila", Desc: "Life is great"},
		&pb.Post{Username: "Nikhila", Desc: "Artic Monkeys#best#ever#music"},
		&pb.Post{Username: "Kavya", Desc: "Music is Life"},
	}
	testwant := new(pb.Posts)
	testwant.PostsList = returnpostsList
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Posts
		wantErr bool
	}{
		// TODO: Add test cases.
		{"GetUserTest1", fields{testPostList, testUserList, testCurrentUser}, args{context.Background(), postsUser}, testwant, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				listOfPosts: tt.fields.listOfPosts,
				listOfUsers: tt.fields.listOfUsers,
				currentUser: tt.fields.currentUser,
			}
			got, err := s.GetPosts(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.GetPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_AddPost(t *testing.T) {
	type fields struct {
		listOfPosts []*pb.Post
		listOfUsers []*pb.User
		currentUser *pb.User
	}
	type args struct {
		ctx context.Context
		in  *pb.Post
	}
	testUserList := []*pb.User{
		&pb.User{Username: "Nikhila", Password: "aGVsbG8=", Followers: []string{"Nikhila", "Kavya"}},
		&pb.User{Username: "Kavya", Password: "eWlwcGVl", Followers: []string{"Kavya", "Navi"}},
		&pb.User{Username: "Navi", Password: "bm9pY2VlZQ==", Followers: []string{"Navi", "Nikhila"}},
	}

	testPostList := []*pb.Post{
		&pb.Post{Username: "Nikhila", Desc: "Life is great"},
		&pb.Post{Username: "Kavya", Desc: "Music is Life"},
		&pb.Post{Username: "Navi", Desc: "Rock and roll all the way"},
		&pb.Post{Username: "Navi", Desc: "Pink floyed-Wish you were here-#rythm#to#ears"},
		&pb.Post{Username: "Nikhila", Desc: "Artic Monkeys#best#ever#music"},
	}
	testCurrentUser := &pb.User{
		Username:  "Nikhila",
		Password:  "aGVsbG8=",
		Followers: []string{"Nikhila", "Kavya"},
	}
	newpost := &pb.Post{
		Username: "Nikhila",
		Desc:     "This is awesome",
	}
	newPostList := []*pb.Post{
		&pb.Post{Username: "Nikhila", Desc: "Life is great"},
		&pb.Post{Username: "Kavya", Desc: "Music is Life"},
		&pb.Post{Username: "Navi", Desc: "Rock and roll all the way"},
		&pb.Post{Username: "Navi", Desc: "Pink floyed-Wish you were here-#rythm#to#ears"},
		&pb.Post{Username: "Nikhila", Desc: "Artic Monkeys#best#ever#music"},
		&pb.Post{Username: "Nikhila", Desc: "This is awesome"},
	}
	testwant := new(pb.Posts)
	testwant.PostsList = newPostList
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Posts
		wantErr bool
	}{
		// TODO: Add test cases.
		{"AddPostTest1", fields{testPostList, testUserList, testCurrentUser}, args{context.Background(), newpost}, testwant, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				listOfPosts: tt.fields.listOfPosts,
				listOfUsers: tt.fields.listOfUsers,
				currentUser: tt.fields.currentUser,
			}
			got, err := s.AddPost(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.AddPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.AddPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_SetCurrentUser(t *testing.T) {
	type fields struct {
		listOfPosts []*pb.Post
		listOfUsers []*pb.User
		currentUser *pb.User
	}
	type args struct {
		ctx context.Context
		in  *pb.User
	}
	testUserList := []*pb.User{
		&pb.User{Username: "Nikhila", Password: "aGVsbG8=", Followers: []string{"Nikhila", "Kavya"}},
		&pb.User{Username: "Kavya", Password: "eWlwcGVl", Followers: []string{"Kavya", "Navi"}},
		&pb.User{Username: "Navi", Password: "bm9pY2VlZQ==", Followers: []string{"Navi", "Nikhila"}},
	}

	testPostList := []*pb.Post{
		&pb.Post{Username: "Nikhila", Desc: "Life is great"},
		&pb.Post{Username: "Kavya", Desc: "Music is Life"},
		&pb.Post{Username: "Navi", Desc: "Rock and roll all the way"},
		&pb.Post{Username: "Navi", Desc: "Pink floyed-Wish you were here-#rythm#to#ears"},
		&pb.Post{Username: "Nikhila", Desc: "Artic Monkeys#best#ever#music"},
	}
	testCurrentUser := &pb.User{
		Username:  "Nikhila",
		Password:  "aGVsbG8=",
		Followers: []string{"Nikhila", "Kavya"},
	}
	currentUser := &pb.User{
		Username:  "Kavya",
		Password:  "eWlwcGVl",
		Followers: []string{"Kavya", "Navi"},
	}
	testwant := new(pb.CurrentUser)
	testwant.CurUser = currentUser
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CurrentUser
		wantErr bool
	}{
		// TODO: Add test cases.
		{"SetCurrentUserTest1", fields{testPostList, testUserList, testCurrentUser}, args{context.Background(), currentUser}, testwant, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				listOfPosts: tt.fields.listOfPosts,
				listOfUsers: tt.fields.listOfUsers,
				currentUser: tt.fields.currentUser,
			}
			got, err := s.SetCurrentUser(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.SetCurrentUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.SetCurrentUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_GetCurrentUser(t *testing.T) {
	type fields struct {
		listOfPosts []*pb.Post
		listOfUsers []*pb.User
		currentUser *pb.User
	}
	type args struct {
		ctx context.Context
		in  *pb.User
	}
	testUserList := []*pb.User{
		&pb.User{Username: "Nikhila", Password: "aGVsbG8=", Followers: []string{"Nikhila", "Kavya"}},
		&pb.User{Username: "Kavya", Password: "eWlwcGVl", Followers: []string{"Kavya", "Navi"}},
		&pb.User{Username: "Navi", Password: "bm9pY2VlZQ==", Followers: []string{"Navi", "Nikhila"}},
	}

	testPostList := []*pb.Post{
		&pb.Post{Username: "Nikhila", Desc: "Life is great"},
		&pb.Post{Username: "Kavya", Desc: "Music is Life"},
		&pb.Post{Username: "Navi", Desc: "Rock and roll all the way"},
		&pb.Post{Username: "Navi", Desc: "Pink floyed-Wish you were here-#rythm#to#ears"},
		&pb.Post{Username: "Nikhila", Desc: "Artic Monkeys#best#ever#music"},
	}
	testCurrentUser := &pb.User{
		Username:  "Nikhila",
		Password:  "aGVsbG8=",
		Followers: []string{"Nikhila", "Kavya"},
	}
	currentUser := &pb.User{
		Username:  "Kavya",
		Password:  "eWlwcGVl",
		Followers: []string{"Kavya", "Navi"},
	}
	testwant := new(pb.CurrentUser)
	testwant.CurUser = currentUser
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CurrentUser
		wantErr bool
	}{
		// TODO: Add test cases.
		{"GetCurrentUserTest1", fields{testPostList, testUserList, testCurrentUser}, args{context.Background(), currentUser}, testwant, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				listOfPosts: tt.fields.listOfPosts,
				listOfUsers: tt.fields.listOfUsers,
				currentUser: tt.fields.currentUser,
			}
			got, err := s.GetCurrentUser(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetCurrentUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.GetCurrentUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_ToggleFollowers(t *testing.T) {
	type fields struct {
		listOfPosts []*pb.Post
		listOfUsers []*pb.User
		currentUser *pb.User
	}
	type args struct {
		ctx context.Context
		in  *pb.FollowUser
	}
	testUserList := []*pb.User{
		&pb.User{Username: "Nikhila", Password: "aGVsbG8=", Followers: []string{"Nikhila", "Kavya"}},
		&pb.User{Username: "Kavya", Password: "eWlwcGVl", Followers: []string{"Kavya", "Navi"}},
		&pb.User{Username: "Navi", Password: "bm9pY2VlZQ==", Followers: []string{"Navi", "Nikhila"}},
	}

	testPostList := []*pb.Post{
		&pb.Post{Username: "Nikhila", Desc: "Life is great"},
		&pb.Post{Username: "Kavya", Desc: "Music is Life"},
		&pb.Post{Username: "Navi", Desc: "Rock and roll all the way"},
		&pb.Post{Username: "Navi", Desc: "Pink floyed-Wish you were here-#rythm#to#ears"},
		&pb.Post{Username: "Nikhila", Desc: "Artic Monkeys#best#ever#music"},
	}
	testCurrentUser := &pb.User{
		Username:  "Nikhila",
		Password:  "aGVsbG8=",
		Followers: []string{"Nikhila", "Kavya"},
	}
	destUser := &pb.User{
		Username: "Navi",
	}
	tempUser1 := &pb.User{
		Username: "Nikhila",
	}
	tempUser2 := &pb.CurrentUser{
		CurUser: tempUser1,
	}
	FUser := &pb.FollowUser{
		SourceUser: tempUser2,
		DestUser:   destUser,
	}
	testwant1 := &pb.User{
		Username:  "Nikhila",
		Followers: []string{"Nikhila", "Kavya", "Navi"},
	}
	testwant2 := &pb.User{
		Username:  "Nikhila",
		Followers: []string{"Nikhila", "Kavya"},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{"ToggleFollowerTest1", fields{testPostList, testUserList, testCurrentUser}, args{context.Background(), FUser}, testwant1, false},
		{"ToggleFollowerTest2", fields{testPostList, testUserList, testCurrentUser}, args{context.Background(), FUser}, testwant2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				listOfPosts: tt.fields.listOfPosts,
				listOfUsers: tt.fields.listOfUsers,
				currentUser: tt.fields.currentUser,
			}
			got, err := s.ToggleFollowers(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.ToggleFollowers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.ToggleFollowers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_GetAllUsers(t *testing.T) {
	type fields struct {
		listOfPosts []*pb.Post
		listOfUsers []*pb.User
		currentUser *pb.User
	}
	type args struct {
		ctx context.Context
		in  *pb.User
	}
	testUserList := []*pb.User{
		&pb.User{Username: "Nikhila", Password: "aGVsbG8=", Followers: []string{"Nikhila", "Kavya"}},
		&pb.User{Username: "Kavya", Password: "eWlwcGVl", Followers: []string{"Kavya", "Navi"}},
		&pb.User{Username: "Navi", Password: "bm9pY2VlZQ==", Followers: []string{"Navi", "Nikhila"}},
	}

	testPostList := []*pb.Post{
		&pb.Post{Username: "Nikhila", Desc: "Life is great"},
		&pb.Post{Username: "Kavya", Desc: "Music is Life"},
		&pb.Post{Username: "Navi", Desc: "Rock and roll all the way"},
		&pb.Post{Username: "Navi", Desc: "Pink floyed-Wish you were here-#rythm#to#ears"},
		&pb.Post{Username: "Nikhila", Desc: "Artic Monkeys#best#ever#music"},
	}
	testCurrentUser := &pb.User{
		Username:  "Nikhila",
		Password:  "aGVsbG8=",
		Followers: []string{"Nikhila", "Kavya"},
	}
	argUser := &pb.User{
		Username: "Nikhila",
	}
	testwant := new(pb.Users)
	testwant.UsersList = testUserList
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Users
		wantErr bool
	}{
		// TODO: Add test cases.
		{"GetAllUsersTest1", fields{testPostList, testUserList, testCurrentUser}, args{context.Background(), argUser}, testwant, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				listOfPosts: tt.fields.listOfPosts,
				listOfUsers: tt.fields.listOfUsers,
				currentUser: tt.fields.currentUser,
			}
			got, err := s.GetAllUsers(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.GetAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

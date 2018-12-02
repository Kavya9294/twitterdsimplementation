package main

import (
	"context"
	"reflect"
	"testing"

	pb "../../authpb"
)

func Test_server_Initialise(t *testing.T) {
	type fields struct {
		listOfPosts []*pb.Post
		listOfUsers []*pb.User
		currentUser *pb.User
	}
	type args struct {
		ctx context.Context
		in  *pb.User
	}
	testPostList := []*pb.Post{
		&pb.Post{Username: "Nikhila", Desc: "Life is great"},
		&pb.Post{Username: "Kavya", Desc: "Music is Life"},
		&pb.Post{Username: "Navi", Desc: "Rock and roll all the way"},
		&pb.Post{Username: "Navi", Desc: "Pink floyed-Wish you were here-#rythm#to#ears"},
		&pb.Post{Username: "Nikhila", Desc: "Artic Monkeys#best#ever#music"},
	}
	testUserList := []*pb.User{
		&pb.User{Username: "Nikhila", Password: "aGVsbG8=", Followers: []string{"Nikhila", "Kavya"}},
		&pb.User{Username: "Kavya", Password: "eWlwcGVl", Followers: []string{"Kavya", "Navi"}},
		&pb.User{Username: "Navi", Password: "bm9pY2VlZQ==", Followers: []string{"Navi", "Nikhila"}},
	}
	testCurrentUser := &pb.User{
		Username: "Nikhila",
		Password: "aGVsbG8=",
		Followers: []string{"Nikhila", "Kavya"},
		
	}
	testpbUser := &pb.User{
		Username: "Nikhila"
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
		{"InitializeTest1",fields{testPostList,testUserList,testCurrentUser},args{context.Background(),testpbUser,nil},}
	}
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
}

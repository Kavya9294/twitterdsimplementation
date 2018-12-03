package auth

import (
	"testing"

	pb "../../web/auth/authpb"
)

func Test_checkUser(t *testing.T) {
	type args struct {
		un       string
		pw       string
		allUsers *pb.Users
	}
	listOfUsers := []*pb.User{
		&pb.User{Username: "Nikhila", Password: "aGVsbG8=", Followers: []string{"Nikhila", "Kavya"}},
		&pb.User{Username: "Kavya", Password: "eWlwcGVl", Followers: []string{"Kavya", "Navi"}},
		&pb.User{Username: "Navi", Password: "bm9pY2VlZQ==", Followers: []string{"Navi", "Nikhila"}},
	}
	user_resp := new(pb.Users)
	user_resp.UsersList = listOfUsers
	tests := []struct {
		name string
		args args
		want bool
	}{

		{"CheckUser Correct Password", args{"Nikhila", "aGVsbG8=", user_resp}, true},
		{"CheckUser Incorrect Password", args{"Nikhila", "aGVsbG8", user_resp}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkUser(tt.args.un, tt.args.pw, tt.args.allUsers); got != tt.want {
				t.Errorf("checkUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

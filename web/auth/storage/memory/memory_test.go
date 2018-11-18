package mymem

import (
	"reflect"
	"testing"
)

func TestAddUser(t *testing.T) {
	Initialize()
	type args struct {
		user_name string
		password  string
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestAddUser1", args{"Poorna", "ak45mnop"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddUser(tt.args.user_name, tt.args.password)
		})
	}
}

/*
func TestUser_GetPosts(t *testing.T) {
	type fields struct {
		Username  string
		Password  string
		Following []string
	}
	type args struct {
		posts []Post
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Post
	}{
		{
			"GetPostsTest1",
			fields{"Kavya", "eWlwcGVl", []string{"Nikhila", "Kavya"}},
			args{PostsList},
			[]Post{
				{"Kavya", "We built this city"},
				{"Nikhila", "Favorite Radio City"},
				{"Nikhila", "Hymn for the Weekend baby"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := User{
				Username:  tt.fields.Username,
				Password:  tt.fields.Password,
				Following: tt.fields.Following,
			}
			if got := user.GetPosts(tt.args.posts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
/*
func TestPost_AppendPost(t *testing.T) {
	Initialize()
	type fields struct {
		Username string
		Desc     string
	}
	tests := []struct {
		name   string
		fields fields
		want   []Post
	}{
		{"AppendPostTest1", fields{"Kavya", "Hello"}, []Post{
			{"Kavya", "We built this city"},
			{"Nikhila", "Favorite Radio City"},
			{"Navi", "On Rock and Roll"},
			{"Navi", "Coldplay rocks"},
			{"Nikhila", "Hymn for the Weekend baby"},
			{"Kavya", "Hello"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			new_post := Post{
				Username: tt.fields.Username,
				Desc:     tt.fields.Desc,
			}
			if got := new_post.AppendPost(); !reflect.DeepEqual(got, tt.want) {

				t.Errorf("Post.AppendPost() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
func TestAddFollower(t *testing.T) {
	Initialize()
	type args struct {
		suser string
		duser string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"AddFollowerTest1", args{"Nikhila", "Navi"}, []string{"Nikhila", "Navi"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddFollower(tt.args.suser, tt.args.duser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFollower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveFollower(t *testing.T) {
	Initialize()
	type args struct {
		suser string
		duser string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"RemoveFollowerTest1", args{"Kavya", "Nikhila"}, []string{"Kavya"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveFollower(tt.args.suser, tt.args.duser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFollower() = %v, want %v", got, tt.want)
			}
		})
	}
}

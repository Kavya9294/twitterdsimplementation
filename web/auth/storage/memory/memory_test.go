package mymem

import (
	"reflect"
	"testing"
)

func TestSetCurrentUser(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"SetCurrentUserTest1", args{"Kavya", "eWlwcGVl"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetCurrentUser(tt.args.username, tt.args.password)
		})
	}
}

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

func TestUser_GetPosts(t *testing.T) {
	Initialize()
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
				{"Nikhila", "Favorite Radio City"},
				{"Nikhila", "Hymn for the Weekend baby"},
				{"Kavya", "We built this city"},
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

func TestAppendPost(t *testing.T) {
	type args struct {
		new_post Post
	}
	tests := []struct {
		name string
		args args
		want []Post
	}{
		// TODO: Add test cases.
		{"AppendPostTest1", args{Post{"Kavya", "Hello"}}, []Post{
			{"Kavya", "We built this city"},
			{"Nikhila", "Favorite Radio City"},
			{"Navi", "On Rock and Roll"},
			{"Navi", "Coldplay rocks"},
			{"Nikhila", "Hymn for the Weekend baby"},
			{"Kavya", "Hello"},
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendPost(tt.args.new_post); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		},
		},
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
func TestToggleFollower(t *testing.T) {
	Initialize()
	type args struct {
		duser string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"ToggleFollowerTest1", args{"Navi"}, []string{"Poorna", "Navi"}},
		{"ToggleFollowerTest2", args{"Poorna"}, []string{"Navi"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToggleFollower(tt.args.duser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToggleFollower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.4
		{"GetAllUsersTest", []string{"Kavya", "Nikhila", "Navi"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllUsers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

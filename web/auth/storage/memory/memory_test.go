package mymem

import (
	"reflect"
	"testing"
)

func TestUser_GetPosts(t *testing.T) {
	posts := PostsList
	type fields struct {
		Username  string
		Password  string
		UserId    int
		Following []int
	}
	type args struct {
		posts *[]Post
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Post
	}{
		{"Test1",
			fields{"Kavya", "yippee", 1, []int{1, 2}}, args{posts},
			[]Post{
				{"Kavya", "We built this city", 1},
				{"Nikhila", "Favorite Radio City", 2},
				{"Navi", "On Rock and Roll", 3},
				{"Navi", "Coldplay rocks", 3},
				{"Nikhila", "Hymn for the Weekend baby", 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				Username:  tt.fields.Username,
				Password:  tt.fields.Password,
				UserId:    tt.fields.UserId,
				Following: tt.fields.Following,
			}
			if got := user.GetPosts(tt.args.posts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPost_AppendPost(t *testing.T) {
	type fields struct {
		Username string
		Desc     string
		UserId   int
	}
	tests := []struct {
		name   string
		fields fields
		want   []Post
	}{
		{"Test2", fields{"Kavya", "Hello", 1}, []Post{
			{"Kavya", "We built this city", 1},
			{"Nikhila", "Favorite Radio City", 2},
			{"Navi", "On Rock and Roll", 3},
			{"Navi", "Coldplay rocks", 3},
			{"Nikhila", "Hymn for the Weekend baby", 2},
			{"Kavya", "Hello", 1},
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			new_post := Post{
				Username: tt.fields.Username,
				Desc:     tt.fields.Desc,
				UserId:   tt.fields.UserId,
			}
			if got := new_post.AppendPost(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Post.AppendPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

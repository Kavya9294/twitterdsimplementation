package mymem

type Post struct {
	Username string
	Desc     string
}

var PostsList []Post

type User struct {
	Username  string
	Password  string
	Following []string
}

var Users []User

var Cur_user User

func Initialize() {
	Users = []User{
		{"Kavya", "eWlwcGVl", []string{"Nikhila", "Kavya"}},
		{"Nikhila", "aGVsbG8=", []string{"Nikhila"}},
		{"Navi", "bm9pY2VlZQ==", []string{"Navi"}},
	}

	PostsList = []Post{
		{"Kavya", "We built this city"},
		{"Nikhila", "Favorite Radio City"},
		{"Navi", "On Rock and Roll"},
		{"Navi", "Coldplay rocks"},
		{"Nikhila", "Hymn for the Weekend baby"},
	}
}

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

func Initialize() {
	Users = []User{
		{"Kavya", "eWlwcGVl", []string{"nikhila", "Kavya"}},
		{"nikhila", "aGVsbG8=", []string{"nikhila"}},
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

package mymem

type Post struct {
	Username string
	Desc     string
}

var PostsList []Post

type User struct {
	Username  string
	Password  string
	Following []int
}

var Users []User

func Initialize() {
	Users = []User{
		{"Kavya", "eWlwcGVl", []int{1, 2}},
		{"nikhila", "aGVsbG8=", []int{2}},
		{"Navi", "bm9pY2VlZQ==", []int{3}},
	}

	PostsList = []Post{
		{"Kavya", "We built this city"},
		{"Nikhila", "Favorite Radio City"},
		{"Navi", "On Rock and Roll"},
		{"Navi", "Coldplay rocks"},
		{"Nikhila", "Hymn for the Weekend baby"},
	}
}

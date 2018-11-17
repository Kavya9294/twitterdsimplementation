package mymem

type Post struct {
	Username string
	Desc     string
	UserId   int
}

var PostsList *[]Post

type User struct {
	Username  string
	Password  string
	UserId    int
	Following []int
}

var Users *[]User

func Initialize() {
	Users = &[]User{
		{"Kavya", "yippee", 1, []int{1, 2}},
		{"Nikhila", "hello", 2, []int{2}},
		{"Navi", "noiceee", 3, []int{3}},
	}

	PostsList = &[]Post{
		{"Kavya", "We built this city", 1},
		{"Nikhila", "Favorite Radio City", 2},
		{"Navi", "On Rock and Roll", 3},
		{"Navi", "Coldplay rocks", 3},
		{"Nikhila", "Hymn for the Weekend baby", 2},
	}
}

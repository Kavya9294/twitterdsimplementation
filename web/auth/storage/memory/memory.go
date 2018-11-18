package mymem

import "fmt"
import "net/http"
import "strings"

func AddUser(user_name string, password string) {
	newuser := user_name
	newPw := password
	var x []string
	x = []string{user_name}
	newUser := User{newuser, newPw, x}
	Users = append(Users, newUser)
}

func (user User) GetPosts(posts []Post) []Post {
	fmt.Print("In get Posts")
	var followingPosts []Post
	all_posts := posts
	for _, following := range user.Following {
		for _, ind_post := range all_posts {
			if ind_post.Username == following {

				followingPosts = append(followingPosts, ind_post)
				fmt.Print("Posts in followingPosts: ", followingPosts)
			}
		}
	}
	return followingPosts
}

func (new_post Post) AppendPost() []Post {
	PostsList = append(PostsList, new_post)
	return PostsList
}

func GetCurrentUser(req *http.Request) User {
	cookie, err := req.Cookie("userInfo")
	if err != nil {
		fmt.Print("Error in getCurrentUser : ", err)
		return User{}
	}
	userInfo := cookie.Value
	fmt.Print("userInfo: ", userInfo)
	temp_string := strings.Split(userInfo, ":")
	un, pw := temp_string[0], temp_string[1]
	for _, user := range Users {
		if un == user.Username && pw == user.Password {
			cur_user := user
			return cur_user
		}
	}

	return User{"", "", []string{""}}

}

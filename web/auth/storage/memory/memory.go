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
	Cur_user = newUser
	fmt.Print("current_user after sign up: ", Cur_user)
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

func AppendPost(new_post Post) []Post {
	PostsList = append(PostsList, new_post)
	return PostsList
}

func SetCurrentUser(username string, password string) {

	Cur_user = User{username, password, []string{username}}
}

func GetCurrentUser(req *http.Request) User {
	cookie, err := req.Cookie("userInfo")
	fmt.Print("Cookie: ", cookie)
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
			Cur_user = user
			return Cur_user
		}
	}
	return User{"", "", []string{""}}

}

func ToggleFollower(duser string) []string {
	pos := -1
	following_list := Cur_user.Following
	var following_new_list []string
	for index, following := range following_list {
		if duser == following {
			pos = index
		}
	}
	if pos == -1 {
		following_list = append(following_list, duser)
		following_new_list = following_list
	} else {
		for i, follow := range following_list {
			if i != pos {
				following_new_list = append(following_new_list, follow)
			}
		}
	}
	Cur_user.Following = following_new_list
	return following_new_list
}

func GetAllUsers() []string {
	var user_list []string
	for _, user := range Users {
		if user.Username != Cur_user.Username {
			user_list = append(user_list, user.Username)
		}
	}

	return user_list
}

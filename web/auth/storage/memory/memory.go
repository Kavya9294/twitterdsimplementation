package mymem

import "fmt"

func AddUser(user_name string, password string) {
	newuser := user_name
	newPw := password
	var x []int
	newUser := User{newuser, newPw, x}
	Users = append(Users, newUser)
}

func (user User) GetPosts(posts []Post) []Post {
	fmt.Print("In get Posts")
	var followingPosts []Post
	//all_posts := posts
	//ub to handle
	/* for _, following := range user.Following {
		for _, ind_post := range all_posts {
			if ind_post.Username == ind_post.Username {

				followingPosts = append(followingPosts, ind_post)
				fmt.Print("Posts in followingPosts: ", followingPosts)
			}
		}
	} */
	return followingPosts
}

func (new_post Post) AppendPost() []Post {
	PostsList = append(PostsList, new_post)
	return PostsList
}

package mymem

import "fmt"

func (user *User) GetPosts(posts *[]Post) []Post {
	fmt.Print("In get Posts")
	var followingPosts []Post
	all_posts := *posts
	for _, following := range user.Following {
		for _, ind_post := range all_posts {
			if ind_post.UserId == following {
				followingPosts = append(followingPosts, ind_post)
				fmt.Print("Posts in followingPosts: ", followingPosts)
			}
		}
	}
	return followingPosts
}

func (new_post *Post) AppendPost() []Post {
	*PostsList = append(*PostsList, *new_post)
	return *PostsList
}

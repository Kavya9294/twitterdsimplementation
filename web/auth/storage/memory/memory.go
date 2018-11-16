package mymem

/* func add_user(user_name string, password string) {
	new_user := user_name
	new_pw := password
	append(user_info, user_name, password)
} */

func check_user(un string, pw string) bool {
	us := make(map[string]user_info)
	i := us[un].password
	status := false
	if i == pw {
		status = true

	}
	return status
}

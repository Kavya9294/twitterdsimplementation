package mymem

import "fmt"

type user_info struct {
	user_name string
	password  string
}

var users []user_info

func basic_user_details() {
	userd := []user_info{
		{"nikhila", "hello"},
		{"ubeee", "yippee"},
		{"poorna", "noiceee"},
	}

	fmt.Print(userd)
}

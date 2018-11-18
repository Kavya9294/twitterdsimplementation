package auth

import (
	"testing"

	"../../web/auth/storage/memory"
)

func Test_checkUser(t *testing.T) {
	mymem.Initialize()
	type args struct {
		un string
		pw string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"CheckUserTest1", args{"Kavya", "eWlwcGVl"}, true},
		{"CheckUserTest1", args{"Kavya", "eWlwcGV"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkUser(tt.args.un, tt.args.pw); got != tt.want {
				t.Errorf("checkUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

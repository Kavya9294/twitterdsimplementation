Project structure:-

We have followed the same structure that was provided by the TA

cmd/web/web.go : All the HTTP request handlers are present here.
cmd/views/* : Contains the templates used for the frontend
web/auth/storage/memory/memory.go : Contians the middleware logic used to access/modify data
web/auth/storage/memory/storage.go : Contains the persistant storage used throughout the system
web/auth/storage/memory/memory_test.go : Contians the unit test cases for the middleware functions used by the rest of the system, excluding HTTP tests.
web/auth/auth.go : Contians the routines related to authentication such as sign up and login handling
web/auth/auth_test.go : Contains the unit test cases for the authentication modules in auth.go, excluding the HTTP tests.

Instructions for usage:-

export GOPATH to the current directory
cd cmd/web
go get ../../web/auth/storage/memory (Importing internal package)
go run web.go (to start the server)

UI:-

Open http://localhost:9090/login
Clear Cookies
Check for existing user : Username ="Nikhila",Password="hello" (to check sign in flow)
Check for new user : Enter username and password in the signup section(to check for sign up flow)

User Flow Description :-

The buttons with the names of the users at the bottom indicate all users that can be followed.
Clicking on them, makes the current user follow that user. Clicking again will lead to unfollow.
After follow, refresh page to see the new following user's posts.
Add text in the post description section and click on Post button to add new post.


Running Test Cases:-

--MIDDLEWARE TESTS--
cd web/auth/storage/memory/
run "go test -v"

--AUTHENTICATION TESTS--
cd web/auth/
run "go test -v"

Known UI bug:-
-> The following button's color consistance. Although, the number of followers is getting updated correctly in the backend,the same is not correctly indicated in the UI.



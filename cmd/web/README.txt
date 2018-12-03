Project structure:-

We have followed the same structure that was provided by the TA

cmd/web/web.go : All the HTTP request handlers are present here.
cmd/views/* : Contains the templates used for the frontend
cmd/memory/memory.go : Contains the gRPC logic used to access/modify data
cmd/memory/memory_test.go : Contains the unit test cases for the gRPC functions used by the rest of the system, excluding HTTP tests.
web/auth/auth.go : Contains the routines related to authentication such as sign up and login handling
web/auth/auth_test.go : Contains the unit test cases for the authentication modules in auth.go, excluding the HTTP tests.

Instructions for usage:-

Open a new terminal window
export GOPATH to the current directory: export GOPATH=$HOME/go
cd web/auth/
Compile the Protobuf file using the following command: protoc -I ./authpb ./authpb/models.proto --go_out=plugins=grpc:authpb
cd ../../cmd/memory/
Start gRPC Server: go run memory.go

Open a new terminal window
cd cmd/web/
Start the HTTP Server(gRPC Client): go run web.go

UI:-

Open http://localhost:9090/
Clear Cookies
Check for existing user : Username ="Nikhila",Password="hello" (to check sign in flow)
Check for new user : Enter username and password in the signup section(to check for sign up flow)

User Flow Description :-

The buttons with the names of the users at the bottom indicate all users that can be followed.
Clicking on them, makes the current user follow that user. Clicking again will lead to unfollow.
After follow, refresh page to see the new following user's posts.
Add text in the post description section and click on Post button to add new post.


Running Test Cases:-

--GRPC TESTS--
cd cmd/memory/
run "go test -v"

--AUTHENTICATION TESTS--
cd web/auth/
run "go test -v"

Known UI bug:-
-> The following button's color consistance. Although, the number of followers is getting updated correctly in the backend,the same is not correctly indicated in the UI.
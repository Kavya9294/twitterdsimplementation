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

Running raft node clusters: 
In order to run the node clusters configured, please goto the $GOPATH(in this case, our project folder) , and execute the following command:
 etcd

Starting a node : This command is used to start the configured node
./bin/goreman run start etcd2

Deleting a node : This command is used to stop the configured node 
./bin/goreman run stop etcd2

To execute the Procfile:
./bin/goreman -f Procfile start

In order to validate the data being persistent on the nodes, please execute the following command:
Navigate to the $GOPATH where etc has been imported.
./src/go.etcd.io/etcd/bin/etcdctl --endpoints= localhost:22379 get "Post"


UI:-

Open http://localhost:9090/
Clear Cookies
Check for new user : Enter username and password in the signup section(to check for sign up flow)
We do not have any test data. Persistence if data is shown as when users are added and their posts are added.

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

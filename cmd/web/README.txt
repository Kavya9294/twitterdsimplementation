Project structure:-

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

Run the cluster like so:
go get go.etcd.io/etcd/etcdctl
go get github.com/mattn/goreman

(1) ./bin/goreman -f Procfile start


Testing for Persistance:-

Starting a node : This command is used to start the configured node
(2) ./bin/goreman run start etcd2

Deleting a node : This command is used to stop the configured node 
(3) ./bin/goreman run stop etcd2

In order to validate the data being persistent on the nodes, please execute the following command:
(4) ./bin/etcdctl --endpoints=localhost:22379 get "Post"

Run the above command once a few posts have been added.
Run it again after stopping and starting node using commands (2) and (3) to check for persistance of data 
even when a node has been brought down and back up again.

UI:-

Open http://localhost:9090/
Clear Cookies
Check for new user : Enter username and password in the signup section(to check for sign up flow)
Login with the same user details at http://localhost:9090/login to test for exisitng user.


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

package main

import (
	"go.etcd.io/etcd/clientv3"

	//"flag"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {

	log.Print("in here")

	cli, err := clientv3.NewFromURL("http://localhost:2379")
	//r := &etcdnaming.GRPCResolver{Client: cli}
	//b := grpc.RoundRobin(r)
	//conn, gerr := grpc.Dial("localhost:8080", grpc.WithBalancer(b))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	if err != nil {
		fmt.Print("error in raft client implementation: ", err)
	}

	defer cli.Close()
	type DemoUser struct {
		Name string
		Age  string
	}
	//var user DemoUser
	user := DemoUser{
		Name: "Kavya",
		Age:  "24",
	}

	type CurrentUser struct {
		Cuser []DemoUser
	}

	var usr CurrentUser
	usr.Cuser = append(usr.Cuser, user)
	usr.Cuser = append(usr.Cuser, user)
	usr.Cuser = append(usr.Cuser, user)

	json_var, _ := json.Marshal(usr)
	_, err = cli.Put(ctx, "foo", string(json_var))
	resp, rerr := cli.Get(ctx, "foo")
	fmt.Print("value")
	if rerr != nil {
		log.Fatal(rerr)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("Value:  %s ", ev.Value)
		var non_json CurrentUser
		_ = json.Unmarshal(ev.Value, &non_json)
		fmt.Print("Json thing\n", non_json)
	}
	cancel()

}

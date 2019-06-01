package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os/exec"

	pb "github.com/JackyChiu/twirpt"
)

func main() {
	out, err := exec.Command("minikube", "service", "twirpt", "--url").Output()
	if err != nil {
		panic(fmt.Errorf("can't find the service's address: %s", err))
	}
	addr := string(bytes.Trim(out, "\n"))

	client := pb.NewHelloWorldProtobufClient(addr, &http.Client{})

	resp, err := client.Hello(context.Background(), &pb.HelloReq{Subject: "World"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Text) // prints "Hello World"
}

package main

import (
	"fmt"
	"net/http"
	"os"

	hello "github.com/JackyChiu/twirpt/internal/hello_world_server"
	pb "github.com/JackyChiu/twirpt/rpc/hello_world"
)

// Run the implementation in a local server
func main() {
	twirpHandler := pb.NewHelloWorldServer(&hello.Server{}, nil)
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	// The generated code includes a method, PathPrefix(), which
	// can be used to mount your service on a mux.
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		panic("no PORT env var set")
	}
	http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
}

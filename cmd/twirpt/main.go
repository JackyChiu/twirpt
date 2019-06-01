package main

import (
	"net/http"

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
	http.ListenAndServe(":8080", mux)
}

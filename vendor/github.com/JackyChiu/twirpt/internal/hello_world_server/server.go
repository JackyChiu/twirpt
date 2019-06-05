package hello_world_server

import (
	"context"

	pb "github.com/JackyChiu/twirpt/rpc/hello_world"
)

type Server struct{}

func (s *Server) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Text: "Hello " + req.Subject}, nil
}

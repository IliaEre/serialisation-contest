package server

import (
	"context"
	pb "proto-docs-service/grpc/docs"
)

type Server struct {
	pb.DocumentServiceServer
}

func GetNewServer() *Server {
	return &Server{}
}

func (s *Server) GetAllByLimitAndOffset(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	
	return nil, nil
}

func (s *Server) Save(ctx context.Context, req *pb.SaveRequest) (*pb.SaveResponse, error) {
	
	return nil, nil
}

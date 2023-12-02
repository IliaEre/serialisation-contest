package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "proto-docs-service/grpc/docs"
	"proto-docs-service/pkg/service"
)

type Server struct {
	pb.DocumentServiceServer

	middle service.ReportService
}

func GetNewServer(middle *service.ReportService) *Server {
	return &Server{middle: *middle}
}

func (s *Server) GetAllByLimitAndOffset(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	docs, err := s.middle.Find(int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(docs) != 0 {
		var typedDocs []*pb.Document
		for _, document := range docs {
			typedDocs = append(typedDocs, document)
		}

		return &pb.GetAllResponse{Documents: typedDocs}, nil
	} else {
		return nil, status.Error(codes.NotFound, "not found")
	}
}

func (s *Server) Save(ctx context.Context, req *pb.SaveRequest) (*pb.SaveResponse, error) {
	err := s.middle.Save(req.GetDocument())
	if err != nil {
		return nil, status.Error(codes.Internal, "ups, problem")
	}
	return &pb.SaveResponse{Message: "ok"}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	err := s.middle.Validate(req.GetDocument())
	if err != nil {
		return nil, status.Error(codes.Internal, "ups, problem")
	}
	return &pb.ValidateResponse{Message: "ok"}, nil
}

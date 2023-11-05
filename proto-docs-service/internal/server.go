package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "proto-docs-service/grpc/docs"
	"proto-docs-service/internal/service"
)

type Server struct {
	pb.DocumentServiceServer
	middle service.ReportService
}

func GetNewServer(middle service.ReportService) *Server {
	return &Server{middle: middle}
}

func (s *Server) GetAllByLimitAndOffset(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	docs := s.middle.Find(int(req.Limit), int(req.Offset))

	if len(docs) != 0 {
		typedDocs := make([]*pb.Document, len(docs))
		for _, document := range docs {
			typedDocs = append(typedDocs, &document)
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

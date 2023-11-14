package service

import "proto-docs-service/grpc/docs"

type ReportServiceInterface interface {
	Save(docs.Document)
	Find(int, int) []docs.Document
}

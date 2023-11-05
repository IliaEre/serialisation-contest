package service

import "proto-docs-service/grpc/docs"

var Docs []docs.Document

type ReportService struct {
	ReportServiceInterface
}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (sv ReportService) Save(doc *docs.Document) error {
	Docs = append(Docs, *doc)
	if len(Docs) > 1000 {
		Docs = Docs[:500] // delete first 500
	}

	return nil
}

func (sv ReportService) Find(limit int, offset int) []docs.Document {
	if offset > len(Docs) {
		return nil
	}
	if offset+limit > len(Docs) {
		return Docs[offset:]
	}

	return Docs[offset : offset+limit]
}

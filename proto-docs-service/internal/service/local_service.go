package service

import "proto-docs-service/internal/model"

type ReportService struct {
	ReportServiceInterface
}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (sv ReportService) Save(doc model.Document) {
	Docs = append(Docs, doc)
	if len(Docs) > 1000 {
		Docs = Docs[:500] // delete first 500
	}
}

func (sv ReportService) Find(limit int, offset int) []model.Document {
	if offset > len(Docs) {
		return nil
	}
	if offset+limit > len(Docs) {
		return Docs[offset:len(Docs)]
	}

	return Docs[offset : offset+limit]
}

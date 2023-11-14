package service

import (
	"flat-docs-service/flat/docs/sample"
)

var Docs []sample.Document

type ReportService struct {
	ReportServiceInterface
}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (sv ReportService) Save(doc sample.Document) error {
	Docs = append(Docs, doc)
	if len(Docs) > 1000 {
		Docs = Docs[:500] // delete first 500
	}
	return nil
}

func (sv ReportService) Find(limit int, offset int) []sample.Document {
	if offset > len(Docs) {
		return nil
	}
	if offset+limit > len(Docs) {
		return Docs[offset:]
	}
	return Docs[offset : offset+limit]
}

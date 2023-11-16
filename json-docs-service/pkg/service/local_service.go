package service

import (
	"json-docs-service/pkg/db"
	"json-docs-service/pkg/model"
)

type ReportService struct {
	ReportServiceInterface

	rc db.ReportClientRepository
}

func NewReportService(mongo db.ReportClientRepository) *ReportService {
	return &ReportService{rc: mongo}
}

func (sv ReportService) Save(doc model.Document) error {
	return sv.rc.Save(doc)
}

func (sv ReportService) Find(limit int, offset int) ([]model.Document, error) {
	return sv.rc.Find(limit, offset)
}

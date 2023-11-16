package service

import (
	"json-docs-service/internal/model"
	"json-docs-service/pkg/db"
)

type ReportService struct {
	ReportServiceInterface
	rc db.ReportMongoRepository
}

func NewReportService(redis db.ReportMongoRepository) *ReportService {
	return &ReportService{rc: redis}
}

func (sv ReportService) Save(doc model.Document) error {
	return sv.rc.Save(doc)
}

func (sv ReportService) Find(limit int, offset int) ([]model.Document, error) {
	return sv.rc.Find(limit, offset)
}

package service

import (
	"proto-docs-service/grpc/docs"
	"proto-docs-service/pkg/db"
)

type ReportService struct {
	ReportServiceInterface
	mr db.ReportMongoRepository
}

func NewReportService(mongo *db.ReportMongoRepository) *ReportService {
	return &ReportService{mr: *mongo}
}

func (sv *ReportService) Save(doc *docs.Document) error {
	err := sv.mr.Save(doc)
	if err != nil {
		return err
	}
	return nil
}

func (sv *ReportService) Find(limit int, offset int) ([]docs.Document, error) {

	return nil, nil
}

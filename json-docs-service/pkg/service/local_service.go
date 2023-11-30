package service

import (
	"errors"
	"json-docs-service/pkg/db"
	"json-docs-service/pkg/model"
	"log"
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

func (sv ReportService) Validate(doc model.Document) error {
	code := doc.Docs.Department.Code
	log.Println("Department code:", code)
	if len(code) > 100 {
		return errors.New("ups, department code so big")
	}
	return nil
}

func (sv ReportService) Find(limit int, offset int) ([]model.Document, error) {
	return sv.rc.Find(limit, offset)
}

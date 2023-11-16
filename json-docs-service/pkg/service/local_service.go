package service

import (
	db "github.com/IliaEre/serialisation-contest/common/db"
	m "github.com/IliaEre/serialisation-contest/common/model"
)

type ReportService struct {
	rc db.ReportMongoRepository
	ReportServiceInterface
}

func NewReportService(mongo *db.ReportMongoRepository) *ReportService {
	return &ReportService{rc: *mongo}
}

func (sv ReportService) Save(doc m.Document) error {
	return sv.rc.Save(doc)
}

func (sv ReportService) Find(limit int, offset int) ([]m.Document, error) {
	return sv.rc.Find(limit, offset)
}

package service

import (
	"errors"
	"log"
	"proto-docs-service/grpc/docs"
	"proto-docs-service/internal/mapper"
	"proto-docs-service/pkg/db"
)

type ReportService struct {
	ReportServiceInterface

	repository db.ReportMongoRepository
	mapper     mapper.ProtoMapper
}

func NewReportService(mongo *db.ReportMongoRepository) *ReportService {
	return &ReportService{repository: *mongo, mapper: *mapper.New()}
}

func (s *ReportService) Save(doc *docs.Document) error {
	model := s.mapper.ProtoToModelDocument(doc)
	return s.repository.Save(model)
}

func (s *ReportService) Find(limit int, offset int) ([]*docs.Document, error) {
	docList, err := s.repository.Find(limit, offset)
	if err != nil {
		log.Println("Error while processing find method:", err)
		return nil, err
	}

	var parsedDocuments []*docs.Document
	for _, v := range docList {
		doc := s.mapper.ModelToProtoDocument(v)
		parsedDocuments = append(parsedDocuments, doc)
	}

	return parsedDocuments, nil
}

func (s *ReportService) Validate(docs *docs.Document) error {
	code := docs.Department.Code
	log.Println("Department code:", code)
	if len(code) > 100 {
		return errors.New("ups, department code so big")
	} else {
		return nil
	}
}

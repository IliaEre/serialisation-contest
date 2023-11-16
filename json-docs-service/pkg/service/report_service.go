package service

import "json-docs-service/internal/model"

type ReportServiceInterface interface {
	Save(model.Document) error
	Find(int, int) ([]model.Document, error)
}

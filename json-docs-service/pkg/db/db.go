package db

import "json-docs-service/internal/model"

type ReportClientRepository interface {
	Save(model.Document) error
	Find(int, int) (model.Document, error)
}

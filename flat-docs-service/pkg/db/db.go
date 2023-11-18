package db

import "flat-docs-service/pkg/model"

type ReportClientRepository interface {
	Save(model.Document) error
	Find(int, int) ([]*model.Document, error)
}

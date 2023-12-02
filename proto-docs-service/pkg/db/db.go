package db

import (
	"proto-docs-service/internal/model"
)

type ReportClientRepository interface {
	Save(*model.Document) error
	Find(int, int) ([]*model.Document, error)
}

package service

import (
	"flat-docs-service/flat/docs/sample"
)

type ReportServiceInterface interface {
	Save(document *sample.Document) error
	Find(int, int) (*[]byte, error)
}

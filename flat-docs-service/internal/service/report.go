package service

import (
	"flat--docs-service/flat/docs/sample"
)

type ReportServiceInterface interface {
	Save(document sample.Document)
	Find(int, int) []sample.Document
}

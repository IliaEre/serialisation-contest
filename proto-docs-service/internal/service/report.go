package service

import "proto-docs-service/internal/model"

type ReportServiceInterface interface {
	Save(model.Document)
	Find(int, int) []model.Document
}

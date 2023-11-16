package service

import m "github.com/IliaEre/serialisation-contest/common/model"

type ReportServiceInterface interface {
	Save(m.Document) error
	Find(int, int) ([]m.Document, error)
}

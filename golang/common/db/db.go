package db

import "github.com/IliaEre/serialisation-contest/golang/common/model"

type ReportClientRepository interface {
	Save(model.Document) error
	Find(int, int) (model.Document, error)
}

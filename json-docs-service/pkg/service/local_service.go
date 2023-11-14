package service

import "json-docs-service/internal/model"

// mock data:
var delivery = model.Delivery{
	Company: "ddeck",
	Address: model.Address{
		Code:      "122",
		Country:   "UK",
		Street:    "Main",
		Apartment: "19",
	},
}

var data = model.Data{
	model.Transaction{
		Type:      "INCOME",
		UUID:      "asd1",
		PointCode: "1231f",
	},
}

var owner = model.Owner{
	UUID:   "1231cdsdfwewq",
	Secret: "12312 dsfd0sfssdf asd",
}

var price = model.Price{
	CategoryA: "1.0",
	CategoryB: "2.0",
	CategoryC: "3.2",
}

var employee = model.Employee{
	Name:    "Ivan",
	Surname: "Popovich",
	Code:    "1231c",
}

var department = model.Department{
	Code:     "123",
	Time:     2312312311,
	Employee: employee,
}

var Docs = []model.Document{
	{
		model.Docs{
			Name:       "test",
			Department: department,
			Price:      price,
			Owner:      owner,
			Data:       data,
			Delivery:   delivery,
		},
	},
}

type ReportService struct {
	ReportServiceInterface
}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (sv ReportService) Save(doc model.Document) {
	Docs = append(Docs, doc)
	if len(Docs) > 1000 {
		Docs = Docs[:500] // delete first 500
	}
}

func (sv ReportService) Find(limit int, offset int) []model.Document {
	if offset > len(Docs) {
		return nil
	}
	if offset+limit > len(Docs) {
		return Docs[offset:]
	}

	return Docs[offset : offset+limit]
}

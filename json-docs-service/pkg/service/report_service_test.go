package service

import (
	"encoding/json"
	"json-docs-service/internal/model"
	"log"
	"testing"
)

// BenchmarkCreateAndMarshal-10    	  168706	      7045 ns/op
func BenchmarkCreateAndMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doc := createDoc()
		_ = doc.Docs.Name // for tests

		bt, err := json.Marshal(doc)
		if err != nil {
			log.Fatal("parse error")
		}

		parsedDoc := new(model.Document)
		if json.Unmarshal(bt, parsedDoc) != nil {
			log.Fatal("parse error")
		}
		_ = parsedDoc.Docs.Name
	}
}

func createDoc() *model.Document {
	return &model.Document{
		Docs: model.Docs{
			Name: "IT",
			Department: model.Department{
				Code: "department_code",
				Time: 303030,
				Employee: model.Employee{
					Name:    "test",
					Surname: "testovich",
					Code:    "code_123",
				},
			},
			Price: model.Price{
				CategoryA: "10.2",
				CategoryB: "10.2",
				CategoryC: "10.2",
			},
			Owner: model.Owner{
				UUID:   "sdfsr132rfds12edsffsdfg",
				Secret: "strange code like uuid but not!",
			},
			Data: model.Data{
				Transaction: model.Transaction{
					Type:      "MOVE",
					UUID:      "sdfsr132rfds12edsffsdfg",
					PointCode: "1231031230",
				},
			},
			Delivery: model.Delivery{
				Company: "NTF_N1",
				Address: model.Address{
					Code:      "A1",
					Country:   "UK",
					Street:    "Main",
					Apartment: "1",
				},
			},
			Goods: getGoods(),
		},
	}
}

func getGoods() []model.Goods {
	goodsList := make([]model.Goods, 1)
	goods := model.Goods{
		Name:   "I_nokla",
		Amount: 1,
		Code:   "1231fdsf1",
	}
	goodsList = append(goodsList, goods)
	return goodsList
}

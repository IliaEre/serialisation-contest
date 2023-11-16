package service

import (
	"encoding/json"
	m "github.com/IliaEre/serialisation-contest/common/model"
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

		parsedDoc := new(m.Document)
		if json.Unmarshal(bt, parsedDoc) != nil {
			log.Fatal("parse error")
		}
		_ = parsedDoc.Docs.Name
	}
}

func createDoc() *m.Document {
	return &m.Document{
		Docs: m.Docs{
			Name: "IT",
			Department: m.Department{
				Code: "department_code",
				Time: 303030,
				Employee: m.Employee{
					Name:    "test",
					Surname: "testovich",
					Code:    "code_123",
				},
			},
			Price: m.Price{
				CategoryA: "10.2",
				CategoryB: "10.2",
				CategoryC: "10.2",
			},
			Owner: m.Owner{
				UUID:   "sdfsr132rfds12edsffsdfg",
				Secret: "strange code like uuid but not!",
			},
			Data: m.Data{
				Transaction: m.Transaction{
					Type:      "MOVE",
					UUID:      "sdfsr132rfds12edsffsdfg",
					PointCode: "1231031230",
				},
			},
			Delivery: m.Delivery{
				Company: "NTF_N1",
				Address: m.Address{
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

func getGoods() []m.Goods {
	goodsList := make([]m.Goods, 1)
	goods := m.Goods{
		Name:   "I_nokla",
		Amount: 1,
		Code:   "1231fdsf1",
	}
	goodsList = append(goodsList, goods)
	return goodsList
}

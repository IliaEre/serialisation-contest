package server

import (
	"google.golang.org/protobuf/proto"
	"log"
	"proto-docs-service/grpc/docs"
	"testing"
)

// BenchmarkCreateAndMarshal-10    	  651063	      1827 ns/op
func BenchmarkCreateAndMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doc := CreateDoc()
		doc.GetName()
		r, e := proto.Marshal(&doc)
		if e != nil {
			log.Fatal("problem with marshal")
		}

		nd := new(docs.Document)
		if proto.Unmarshal(r, nd) != nil {
			log.Fatal("problem with unmarshal")
		}
		nd.GetName()
	}
}

// BenchmarkCreate-10    	 2265710	       530.6 ns/op
func BenchmarkCreate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := CreateDoc()
		d.GetName()
	}
}

func CreateDoc() docs.Document {
	return docs.Document{
		Name:       "IT",
		Department: createDepartment(),
		Price:      createPrice(),
		Owner:      createOwner(),
		Data:       createData(),
		Delivery:   createDelivery(),
		Goods:      createGoods(),
	}
}

func createDepartment() *docs.Department {
	return &docs.Department{
		Code:     "department_code",
		Time:     303030,
		Employee: createEmployee(),
	}
}

func createEmployee() *docs.Employee {
	return &docs.Employee{
		Name:    "test",
		Surname: "testovich",
		Code:    "code_123",
	}
}

func createPrice() *docs.Price {
	return &docs.Price{
		CategoryA: "10.2",
		CategoryB: "10.2",
		CategoryC: "10.2",
	}
}

func createOwner() *docs.Owner {
	return &docs.Owner{
		Uuid:   "sdfsr132rfds12edsffsdfg",
		Secret: "strange code like uuid but not!",
	}
}

func createData() *docs.Data {
	return &docs.Data{Transaction: createTransaction()}
}

func createTransaction() *docs.Transaction {
	return &docs.Transaction{
		Type:      "MOVE",
		Uuid:      "sdfsr132rfds12edsffsdfg",
		PointCode: "1231031230",
	}
}

func createDelivery() *docs.Delivery {
	return &docs.Delivery{
		Company: "NTF_N1",
		Address: buildAddress(),
	}
}

func buildAddress() *docs.Address {
	return &docs.Address{
		Code:      "A1",
		Country:   "UK",
		Street:    "Main",
		Apartment: "1",
	}
}

func createGoods() []*docs.Goods {
	goodsList := make([]*docs.Goods, 0)
	doc := &docs.Goods{
		Name:   "I_nokla",
		Code:   "1231fdsf1",
		Amount: 10,
	}
	goodsList = append(goodsList, doc)
	return goodsList
}

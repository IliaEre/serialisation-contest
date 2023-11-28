package main

import (
	uuid "github.com/satori/go.uuid"
	pb "grpc-load-test/docs"
)

func createDoc(name string) *pb.Document {
	return &pb.Document{
		Name:       name + uuid.NewV4().String(),
		Department: createDepartment(),
		Price:      createPrice(),
		Owner:      createOwner(),
		Data:       createData(),
		Delivery:   createDelivery(),
		Goods:      createGoods(),
	}
}

func createDepartment() *pb.Department {
	return &pb.Department{
		Code:     "department_code",
		Time:     303030,
		Employee: createEmployee(),
	}
}

func createEmployee() *pb.Employee {
	return &pb.Employee{
		Name:    "test",
		Surname: "testovich",
		Code:    "code_123",
	}
}

func createPrice() *pb.Price {
	return &pb.Price{
		CategoryA: "10.2",
		CategoryB: "10.2",
		CategoryC: "10.2",
	}
}

func createOwner() *pb.Owner {
	return &pb.Owner{
		Uuid:   "sdfsr132rfds12edsffsdfg",
		Secret: "strange code like uuid but not!",
	}
}

func createData() *pb.Data {
	return &pb.Data{Transaction: createTransaction()}
}

func createTransaction() *pb.Transaction {
	return &pb.Transaction{
		Type:      "MOVE",
		Uuid:      "sdfsr132rfds12edsffsdfg",
		PointCode: "1231031230",
	}
}

func createDelivery() *pb.Delivery {
	return &pb.Delivery{
		Company: "NTF_N1",
		Address: buildAddress(),
	}
}

func buildAddress() *pb.Address {
	return &pb.Address{
		Code:      "A1",
		Country:   "UK",
		Street:    "Main",
		Apartment: "1",
	}
}

func createGoods() []*pb.Goods {
	goodsList := make([]*pb.Goods, 0)
	doc := &pb.Goods{
		Name:   "I_nokla",
		Code:   "1231fdsf1",
		Amount: 10,
	}
	goodsList = append(goodsList, doc)
	return goodsList
}

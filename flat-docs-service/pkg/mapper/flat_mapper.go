package mapper

import (
	"flat-docs-service/flat/docs/sample"
	flatbuffers "github.com/google/flatbuffers/go"
)

func CreateDocument(builder *flatbuffers.Builder, doc *sample.Document) flatbuffers.UOffsetT {
	name := builder.CreateString(string(doc.Name()))

	department := CreateDepartment(builder, doc.Department(new(sample.Department)))
	price := CreatePrice(builder, doc.Price(new(sample.Price)))
	owner := CreateOwner(builder, doc.Owner(new(sample.Owner)))
	data := CreateData(builder, doc.Data(new(sample.Data)))
	delivery := CreateDelivery(builder, doc.Delivery(new(sample.Delivery)))

	goodsObj := new(sample.Goods)
	goodsList := make([]flatbuffers.UOffsetT, doc.GoodsLength())
	for i := 0; i < doc.GoodsLength(); i++ {
		doc.Goods(goodsObj, i)
		cg := CreateGoods(builder, goodsObj)
		goodsList = append(goodsList, cg)
	}

	sample.DocumentStartGoodsVector(builder, len(goodsList))
	for i := len(goodsList) - 1; i >= 0; i-- {
		builder.PrependUOffsetT(goodsList[i])
	}
	goodsVector := builder.EndVector(len(goodsList))

	sample.DocumentStart(builder)
	sample.DocumentAddName(builder, name)
	sample.DocumentAddDepartment(builder, department)
	sample.DocumentAddPrice(builder, price)
	sample.DocumentAddOwner(builder, owner)
	sample.DocumentAddData(builder, data)
	sample.DocumentAddDelivery(builder, delivery)
	sample.DocumentAddGoods(builder, goodsVector)
	return sample.DocumentEnd(builder)
}

func CreateDepartment(builder *flatbuffers.Builder, dep *sample.Department) flatbuffers.UOffsetT {
	code := builder.CreateString(string(dep.Code()))
	employee := CreateEmployee(builder, dep.Employee(new(sample.Employee)))
	sample.DepartmentStart(builder)
	sample.DepartmentAddCode(builder, code)
	sample.DepartmentAddTime(builder, dep.Time())
	sample.DepartmentAddEmployee(builder, employee)
	return sample.DepartmentEnd(builder)
}

func CreateEmployee(builder *flatbuffers.Builder, emp *sample.Employee) flatbuffers.UOffsetT {
	name := builder.CreateString(string(emp.Name()))
	surname := builder.CreateString(string(emp.Surname()))
	code := builder.CreateString(string(emp.Code()))

	sample.EmployeeStart(builder)
	sample.EmployeeAddName(builder, name)
	sample.EmployeeAddSurname(builder, surname)
	sample.EmployeeAddCode(builder, code)
	return sample.EmployeeEnd(builder)
}

func CreatePrice(builder *flatbuffers.Builder, price *sample.Price) flatbuffers.UOffsetT {
	categoryA := builder.CreateString(string(price.CategoryA()))
	categoryB := builder.CreateString(string(price.CategoryB()))
	categoryC := builder.CreateString(string(price.CategoryC()))

	sample.PriceStart(builder)
	sample.PriceAddCategoryA(builder, categoryA)
	sample.PriceAddCategoryB(builder, categoryB)
	sample.PriceAddCategoryC(builder, categoryC)
	return sample.PriceEnd(builder)
}

func CreateOwner(builder *flatbuffers.Builder, owner *sample.Owner) flatbuffers.UOffsetT {
	uuid := builder.CreateString(string(owner.Uuid()))
	secret := builder.CreateString(string(owner.Secret()))

	sample.OwnerStart(builder)
	sample.OwnerAddUuid(builder, uuid)
	sample.OwnerAddSecret(builder, secret)
	return sample.OwnerEnd(builder)
}

func CreateData(builder *flatbuffers.Builder, data *sample.Data) flatbuffers.UOffsetT {
	transaction := CreateTransaction(builder, data.Transaction(new(sample.Transaction)))

	sample.DataStart(builder)
	sample.DataAddTransaction(builder, transaction)
	return sample.DataEnd(builder)
}

func CreateTransaction(builder *flatbuffers.Builder, transaction *sample.Transaction) flatbuffers.UOffsetT {
	transType := builder.CreateString(string(transaction.Type()))
	uuid := builder.CreateString(string(transaction.Uuid()))
	pointCode := builder.CreateString(string(transaction.PointCode()))

	sample.TransactionStart(builder)
	sample.TransactionAddType(builder, transType)
	sample.TransactionAddUuid(builder, uuid)
	sample.TransactionAddPointCode(builder, pointCode)
	return sample.TransactionEnd(builder)
}

func CreateDelivery(builder *flatbuffers.Builder, delivery *sample.Delivery) flatbuffers.UOffsetT {
	company := builder.CreateString(string(delivery.Company()))
	address := CreateAddress(builder, delivery.Address(new(sample.Address)))

	sample.DeliveryStart(builder)
	sample.DeliveryAddCompany(builder, company)
	sample.DeliveryAddAddress(builder, address)
	return sample.DeliveryEnd(builder)
}

func CreateAddress(builder *flatbuffers.Builder, address *sample.Address) flatbuffers.UOffsetT {
	code := builder.CreateString(string(address.Code()))
	country := builder.CreateString(string(address.Country()))
	street := builder.CreateString(string(address.Street()))
	apartment := builder.CreateString(string(address.Apartment()))

	sample.AddressStart(builder)
	sample.AddressAddCode(builder, code)
	sample.AddressAddCountry(builder, country)
	sample.AddressAddStreet(builder, street)
	sample.AddressAddApartment(builder, apartment)
	return sample.AddressEnd(builder)
}

func CreateGoods(builder *flatbuffers.Builder, goods *sample.Goods) flatbuffers.UOffsetT {
	name := builder.CreateString(string(goods.Name()))
	code := builder.CreateString(string(goods.Code()))

	sample.GoodsStart(builder)
	sample.GoodsAddName(builder, name)
	sample.GoodsAddAmount(builder, goods.Amount())
	sample.GoodsAddCode(builder, code)
	g := sample.GoodsEnd(builder)
	return g
}

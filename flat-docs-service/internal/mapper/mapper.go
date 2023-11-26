package mapper

import (
	"flat-docs-service/flat/docs/sample"
	"flat-docs-service/pkg/model"
	flatbuffers "github.com/google/flatbuffers/go"
)

type FlatMapper struct {
}

func New() *FlatMapper {
	return &FlatMapper{}
}

func (pm *FlatMapper) MapToModel(flatDoc *sample.Document) *model.Document {
	var goods []model.Goods
	flatGoods := new(sample.Goods)
	for i := 0; i < flatDoc.GoodsLength(); i++ {
		if flatDoc.Goods(flatGoods, i) {
			product := model.Goods{
				Name:   string(flatGoods.Name()),
				Amount: int(flatGoods.Amount()),
				Code:   string(flatGoods.Code()),
			}
			goods = append(goods, product)
		}
	}

	flatDepartment := new(sample.Department)
	flatDoc.Department(flatDepartment)

	flatEmployee := new(sample.Employee)
	flatDepartment.Employee(flatEmployee)

	department := model.Department{
		Code: string(flatDepartment.Code()),
		Time: flatDepartment.Time(),
		Employee: model.Employee{
			Name:    string(flatEmployee.Name()),
			Surname: string(flatEmployee.Surname()),
			Code:    string(flatEmployee.Code()),
		},
	}

	flatPrice := new(sample.Price)
	flatDoc.Price(flatPrice)

	flatOwner := new(sample.Owner)
	flatDoc.Owner(flatOwner)

	flatData := new(sample.Data)
	flatDoc.Data(flatData)
	flatTrans := new(sample.Transaction)
	flatData.Transaction(flatTrans)

	flatDelivery := new(sample.Delivery)
	flatDoc.Delivery(flatDelivery)
	flatAddress := new(sample.Address)
	flatDelivery.Address(flatAddress)

	return &model.Document{
		Docs: model.Docs{
			Name:       string(flatDoc.Name()),
			Department: department,
			Price: model.Price{
				CategoryA: string(flatPrice.CategoryA()),
				CategoryB: string(flatPrice.CategoryB()),
				CategoryC: string(flatPrice.CategoryC()),
			},
			Owner: model.Owner{
				UUID:   string(flatOwner.Uuid()),
				Secret: string(flatOwner.Secret()),
			},
			Data: model.Data{
				Transaction: model.Transaction{
					Type:      string(flatTrans.Type()),
					UUID:      string(flatTrans.Uuid()),
					PointCode: string(flatTrans.PointCode()),
				},
			},
			Delivery: model.Delivery{
				Company: string(flatDelivery.Company()),
				Address: model.Address{
					Code:      string(flatAddress.Code()),
					Country:   string(flatAddress.Country()),
					Street:    string(flatAddress.Street()),
					Apartment: string(flatAddress.Apartment()),
				},
			},
			Goods: goods,
		},
	}
}

func (pm *FlatMapper) MapToFlat(model *model.Document, builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	name := builder.CreateString(model.Docs.Name)
	department := createDepartment(builder, &model.Docs.Department)
	price := createPrice(builder, &model.Docs.Price)
	owner := createOwner(builder, &model.Docs.Owner)
	data := createData(builder, &model.Docs.Data)
	delivery := createDelivery(builder, &model.Docs.Delivery)

	goodsList := make([]flatbuffers.UOffsetT, 0)
	for i := 0; i < len(model.Docs.Goods); i++ {
		cg := createGoods(builder, &model.Docs.Goods[i])
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

func createDepartment(builder *flatbuffers.Builder, department *model.Department) flatbuffers.UOffsetT {
	code := builder.CreateString(department.Code)
	employee := createEmployee(builder, &department.Employee)
	sample.DepartmentStart(builder)
	sample.DepartmentAddCode(builder, code)
	sample.DepartmentAddTime(builder, department.Time)
	sample.DepartmentAddEmployee(builder, employee)
	return sample.DepartmentEnd(builder)
}

func createEmployee(builder *flatbuffers.Builder, employee *model.Employee) flatbuffers.UOffsetT {
	name := builder.CreateString(employee.Name)
	surname := builder.CreateString(employee.Surname)
	code := builder.CreateString(employee.Code)
	sample.EmployeeStart(builder)
	sample.EmployeeAddName(builder, name)
	sample.EmployeeAddSurname(builder, surname)
	sample.EmployeeAddCode(builder, code)
	return sample.EmployeeEnd(builder)
}

func createPrice(builder *flatbuffers.Builder, price *model.Price) flatbuffers.UOffsetT {
	categoryA := builder.CreateString(price.CategoryA)
	categoryB := builder.CreateString(price.CategoryB)
	categoryC := builder.CreateString(price.CategoryC)
	sample.PriceStart(builder)
	sample.PriceAddCategoryA(builder, categoryA)
	sample.PriceAddCategoryB(builder, categoryB)
	sample.PriceAddCategoryC(builder, categoryC)
	return sample.PriceEnd(builder)
}

func createOwner(builder *flatbuffers.Builder, owner *model.Owner) flatbuffers.UOffsetT {
	uuid := builder.CreateString(owner.UUID)
	secret := builder.CreateString(owner.Secret)
	sample.OwnerStart(builder)
	sample.OwnerAddUuid(builder, uuid)
	sample.OwnerAddSecret(builder, secret)
	return sample.OwnerEnd(builder)
}

func createData(builder *flatbuffers.Builder, data *model.Data) flatbuffers.UOffsetT {
	transaction := createTransaction(builder, &data.Transaction)
	sample.DataStart(builder)
	sample.DataAddTransaction(builder, transaction)
	return sample.DataEnd(builder)
}

func createTransaction(builder *flatbuffers.Builder, transaction *model.Transaction) flatbuffers.UOffsetT {
	transType := builder.CreateString(transaction.Type)
	uuid := builder.CreateString(transaction.UUID)
	pointCode := builder.CreateString(transaction.PointCode)
	sample.TransactionStart(builder)
	sample.TransactionAddType(builder, transType)
	sample.TransactionAddUuid(builder, uuid)
	sample.TransactionAddPointCode(builder, pointCode)
	return sample.TransactionEnd(builder)
}

func createDelivery(builder *flatbuffers.Builder, delivery *model.Delivery) flatbuffers.UOffsetT {
	company := builder.CreateString(delivery.Company)
	address := createAddress(builder, &delivery.Address)
	sample.DeliveryStart(builder)
	sample.DeliveryAddCompany(builder, company)
	sample.DeliveryAddAddress(builder, address)
	return sample.DeliveryEnd(builder)
}

func createAddress(builder *flatbuffers.Builder, address *model.Address) flatbuffers.UOffsetT {
	code := builder.CreateString(address.Code)
	country := builder.CreateString(address.Country)
	street := builder.CreateString(address.Street)
	apartment := builder.CreateString(address.Apartment)
	sample.AddressStart(builder)
	sample.AddressAddCode(builder, code)
	sample.AddressAddCountry(builder, country)
	sample.AddressAddStreet(builder, street)
	sample.AddressAddApartment(builder, apartment)
	return sample.AddressEnd(builder)
}

func createGoods(builder *flatbuffers.Builder, goods *model.Goods) flatbuffers.UOffsetT {
	name := builder.CreateString(goods.Name)
	code := builder.CreateString(goods.Code)
	sample.GoodsStart(builder)
	sample.GoodsAddName(builder, name)
	sample.GoodsAddAmount(builder, int32(goods.Amount))
	sample.GoodsAddCode(builder, code)
	g := sample.GoodsEnd(builder)
	return g
}

package service

import (
	"flat--docs-service/flat/docs/sample"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	BenchCall  = 100_000
	BatchSize  = 1000
	bufferSize = 1024
)

func BenchmarkCreateAndCheck(b *testing.B) {
	b.N = BenchCall

	for i := 0; i < BenchCall; i++ {
		buf := BuildDocs()
		parserDocs := sample.GetRootAsDocument(buf, 0)

		assert.Equal(b, "IT", string(parserDocs.Name()))
		goodsName := "I_nokla"
		parsedGoods := new(sample.Goods)
		for i := 0; i < parserDocs.GoodsLength(); i++ {
			if parserDocs.Goods(parsedGoods, i) {
				assert.Equal(b, goodsName, string(parsedGoods.Name()))
			}
		}
	}
}

func BenchmarkCreate(b *testing.B) {
	b.N = BenchCall

	for i := 0; i < BenchCall; i++ {
		buf := BuildDocs()
		parserDocs := sample.GetRootAsDocument(buf, 0)
		assert.Equal(b, "IT", string(parserDocs.Name()))
	}
}

func BuildDocs() []byte {
	builder := flatbuffers.NewBuilder(bufferSize)
	// build internal objects
	employee := BuildEmployee(builder)
	department := BuildDepartment(builder, employee)
	prices := BuildPrices(builder)
	owner := BuildOwner(builder)
	transaction := BuildTransaction(builder)
	data := BuildDate(builder, transaction)
	address := BuildAddress(builder)
	delivery := BuildDelivery(builder, address)
	goods := BuildGoods(builder)

	sample.DocumentStartGoodsVector(builder, 1)
	builder.PrependUOffsetT(goods)
	goodsVector := builder.EndVector(1)
	// fill doc's name:
	documentCode := builder.CreateString("IT")
	// build doc
	sample.DocumentStart(builder)
	sample.DocumentAddName(builder, documentCode)
	sample.DocumentAddDepartment(builder, department)
	sample.DocumentAddPrice(builder, prices)
	sample.DocumentAddOwner(builder, owner)
	sample.DocumentAddData(builder, data)
	sample.DocumentAddDelivery(builder, delivery)
	sample.DocumentAddGoods(builder, goodsVector)

	docs := sample.DocumentEnd(builder)
	builder.Finish(docs)
	buf := builder.FinishedBytes()
	return buf
}

func BuildEmployee(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	// create employee's fields
	name := builder.CreateString("test")
	surname := builder.CreateString("testovich")
	code := builder.CreateString("code_123")
	// build employee
	sample.EmployeeStart(builder)
	sample.EmployeeAddName(builder, name)
	sample.EmployeeAddSurname(builder, surname)
	sample.EmployeeAddCode(builder, code)
	employee := sample.EmployeeEnd(builder)
	return employee
}

func BuildDepartment(builder *flatbuffers.Builder, employee flatbuffers.UOffsetT) flatbuffers.UOffsetT {
	// fill department's fields
	departmentCode := builder.CreateString("department_code")
	// build department
	sample.DepartmentStart(builder)
	sample.DepartmentAddCode(builder, departmentCode)
	sample.DepartmentAddTime(builder, 303030)
	sample.DepartmentAddEmployee(builder, employee)
	department := sample.DepartmentEnd(builder)
	return department
}

func BuildPrices(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	// create prices
	priceA := builder.CreateString("10.2")
	priceB := builder.CreateString("10.2")
	priceC := builder.CreateString("10.2")
	// build price
	sample.PriceStart(builder)
	sample.PriceAddCategoryA(builder, priceA)
	sample.PriceAddCategoryB(builder, priceB)
	sample.PriceAddCategoryC(builder, priceC)
	prices := sample.PriceEnd(builder)
	return prices
}

func BuildOwner(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	// create fields
	uuid := builder.CreateString("sdfsr132rfds12edsffsdfg")
	secret := builder.CreateString("strange code like uuid but not!")
	sample.OwnerStart(builder)
	sample.OwnerAddUuid(builder, uuid)
	sample.OwnerAddSecret(builder, secret)
	owner := sample.OwnerEnd(builder)
	return owner
}

func BuildDate(builder *flatbuffers.Builder, transaction flatbuffers.UOffsetT) flatbuffers.UOffsetT {
	sample.DataStart(builder)
	sample.DataAddTransaction(builder, transaction)
	data := sample.DataEnd(builder)
	return data
}

func BuildTransaction(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	tType := builder.CreateString("MOVE")
	uuid := builder.CreateString("sdfsr132rfds12edsffsdfg")
	pointCode := builder.CreateString("1231031230")
	sample.TransactionStart(builder)
	sample.TransactionAddType(builder, tType)
	sample.TransactionAddUuid(builder, uuid)
	sample.TransactionAddPointCode(builder, pointCode)
	transaction := sample.TransactionEnd(builder)
	return transaction
}

func BuildDelivery(builder *flatbuffers.Builder, address flatbuffers.UOffsetT) flatbuffers.UOffsetT {
	company := builder.CreateString("NTF_N1")
	sample.DeliveryStart(builder)
	sample.DeliveryAddCompany(builder, company)
	sample.DeliveryAddAddress(builder, address)
	delivery := sample.DeliveryEnd(builder)
	return delivery
}

func BuildAddress(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	code := builder.CreateString("A1")
	country := builder.CreateString("UK")
	street := builder.CreateString("Main")
	apartment := builder.CreateString("1")

	sample.AddressStart(builder)
	sample.AddressAddCode(builder, code)
	sample.AddressAddCountry(builder, country)
	sample.AddressAddStreet(builder, street)
	sample.AddressAddApartment(builder, apartment)
	address := sample.AddressEnd(builder)
	return address
}

func BuildGoods(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	name := builder.CreateString("I_nokla")
	code := builder.CreateString("1231fdsf1")

	sample.GoodsStart(builder)
	sample.GoodsAddName(builder, name)
	sample.GoodsAddAmount(builder, 10)
	sample.GoodsAddCode(builder, code)
	goods := sample.GoodsEnd(builder)
	return goods
}

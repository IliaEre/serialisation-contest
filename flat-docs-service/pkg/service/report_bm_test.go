package service

import (
	"flat-docs-service/flat/docs/sample"
	"flat-docs-service/internal/builder"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	poolCapacity = 100
)

func BenchmarkCreateAndCheckOneField(b *testing.B) {
	bb := builder.NewBuilderPool(poolCapacity)

	for i := 0; i < b.N; i++ {
		newBuilder := bb.Get()
		buf := BuildDocs(newBuilder)
		parserDocs := sample.GetRootAsDocument(buf, 0)

		bb.Put(newBuilder)

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

// BenchmarkCreatePool-10    	 1709313	       698.6 ns/op
func BenchmarkCreatePool(b *testing.B) {
	builderPool := builder.NewBuilderPool(10)

	for i := 0; i < b.N; i++ {
		b := builderPool.Get()
		buf := BuildDocs(b)
		doc := sample.GetRootAsDocument(buf, 0)
		_ = doc.Name()

		builderPool.Put(b)
	}
}

// BenchmarkCreateAndMarshalBuilderPool-10    	  344282	      3467 ns/op
func BenchmarkCreateAndMarshalBuilderPool(b *testing.B) {
	builderPool := builder.NewBuilderPool(10)

	for i := 0; i < b.N; i++ {
		currentBuilder := builderPool.Get()

		buf := BuildDocs(currentBuilder)
		doc := sample.GetRootAsDocument(buf, 0)
		_ = doc.Name()

		sb := doc.Table().Bytes
		cd := sample.GetRootAsDocument(sb, 0)
		_ = cd.Name()

		builderPool.Put(currentBuilder)
	}
}

func BuildDocs(builder *flatbuffers.Builder) []byte {
	// build internal objects
	department := buildDepartment(builder)
	prices := buildPrices(builder)
	owner := buildOwner(builder)
	data := buildData(builder)
	delivery := buildDelivery(builder)

	goods := make([]flatbuffers.UOffsetT, 1)
	for i := 0; i < 1; i++ {
		goods[i] = BuildGoods(builder)
	}

	sample.DocumentStartGoodsVector(builder, 1)
	builder.PrependUOffsetT(goods[0])
	goodsVector := builder.EndVector(1)

	// build doc
	sample.DocumentStart(builder)
	sample.DocumentAddName(builder, builder.CreateString("IT"))
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

func buildEmployee(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	name := builder.CreateString("test")
	surname := builder.CreateString("testovich")
	code := builder.CreateString("code_123")
	sample.EmployeeStart(builder)
	sample.EmployeeAddName(builder, name)
	sample.EmployeeAddSurname(builder, surname)
	sample.EmployeeAddCode(builder, code)
	employee := sample.EmployeeEnd(builder)
	return employee
}

func buildDepartment(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	departmentCode := builder.CreateString("department_code")
	employee := buildEmployee(builder)

	sample.DepartmentStart(builder)
	sample.DepartmentAddCode(builder, departmentCode)
	sample.DepartmentAddTime(builder, 303030)
	sample.DepartmentAddEmployee(builder, employee)
	department := sample.DepartmentEnd(builder)
	return department
}

func buildPrices(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	priceA := builder.CreateString("10.2")
	priceB := builder.CreateString("10.3")
	priceC := builder.CreateString("10.4")
	sample.PriceStart(builder)
	sample.PriceAddCategoryA(builder, priceA)
	sample.PriceAddCategoryB(builder, priceB)
	sample.PriceAddCategoryC(builder, priceC)
	prices := sample.PriceEnd(builder)
	return prices
}

func buildOwner(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	uuid := builder.CreateString("sdfsr132rfds12edsffsdfg")
	secret := builder.CreateString("strange code like uuid but not!")
	sample.OwnerStart(builder)
	sample.OwnerAddUuid(builder, uuid)
	sample.OwnerAddSecret(builder, secret)
	owner := sample.OwnerEnd(builder)
	return owner
}

func buildData(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	transaction := buildTransaction(builder)

	sample.DataStart(builder)
	sample.DataAddTransaction(builder, transaction)
	data := sample.DataEnd(builder)
	return data
}

func buildTransaction(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
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

func buildDelivery(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	address := buildAddress(builder)
	company := builder.CreateString("NTF_N1")

	sample.DeliveryStart(builder)
	sample.DeliveryAddCompany(builder, company)
	sample.DeliveryAddAddress(builder, address)
	delivery := sample.DeliveryEnd(builder)
	return delivery
}

func buildAddress(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
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

func buildGoodsVector(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	goods := make([]flatbuffers.UOffsetT, 1)
	for i := 0; i < 1; i++ {
		goods[i] = BuildGoods(builder)
	}

	sample.DocumentStartGoodsVector(builder, 1)
	builder.PrependUOffsetT(goods[0])
	goodsVector := builder.EndVector(1)
	return goodsVector
}

package service

import (
	"bufio"
	"flat--docs-service/flat/docs/sample"
	flatbuffers "github.com/google/flatbuffers/go"
	"os"
)

func GenerateFind() {
	builder := flatbuffers.NewBuilder(bufferSize)
	sample.FindRequestStart(builder)
	sample.FindRequestAddLimit(builder, 10)
	sample.FindRequestAddOffset(builder, 0)
	response := sample.FindRequestEnd(builder)
	builder.Finish(response)
	bytes := builder.FinishedBytes()

	f, err := os.Create("request.binary")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.Write(bytes)
	check(err)
	w.Flush()
}

func GenerateSave() {
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

	f, err := os.Create("save_request.binary")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.Write(buf)
	check(err)
	w.Flush()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

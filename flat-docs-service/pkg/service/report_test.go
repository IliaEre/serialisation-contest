package service

import (
	"bufio"
	"flat-docs-service/flat/docs/sample"
	"flat-docs-service/internal/builder"
	"fmt"
	"os"
	"testing"
)

func TestShouldGenerateFind(t *testing.T) {
	bp := builder.NewBuilderPool(1)
	bb := bp.Get()
	defer bp.Put(bb)

	sample.FindRequestStart(bb)
	sample.FindRequestAddLimit(bb, 10)
	sample.FindRequestAddOffset(bb, 0)
	response := sample.FindRequestEnd(bb)
	bb.Finish(response)
	bytes := bb.FinishedBytes()

	f, err := os.Create("request.binary")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.Write(bytes)
	check(err)
	w.Flush()
}

func TestShouldGenerateSave(t *testing.T) {
	bp := builder.NewBuilderPool(1)
	bb := bp.Get()
	defer bp.Put(bb)

	// build internal objects
	department := buildDepartment(bb)
	prices := buildPrices(bb)
	owner := buildOwner(bb)
	data := buildData(bb)
	delivery := buildDelivery(bb)
	goods := buildGoodsVector(bb)

	// fill doc's name:
	documentCode := bb.CreateString("IT")
	// build doc
	sample.DocumentStart(bb)
	sample.DocumentAddName(bb, documentCode)
	sample.DocumentAddDepartment(bb, department)
	sample.DocumentAddPrice(bb, prices)
	sample.DocumentAddOwner(bb, owner)
	sample.DocumentAddData(bb, data)
	sample.DocumentAddDelivery(bb, delivery)
	sample.DocumentAddGoods(bb, goods)

	docs := sample.DocumentEnd(bb)

	sample.SaveRequestStart(bb)
	sample.SaveRequestAddDocument(bb, docs)
	se := sample.SaveRequestEnd(bb)
	bb.Finish(se)

	rb := bb.FinishedBytes()

	err := os.WriteFile("save.bin", rb, 0644)
	check(err)
	fmt.Println("Structure written into file successfully")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

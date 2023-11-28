package service

import (
	"errors"
	"flat-docs-service/flat/docs/sample"
	"flat-docs-service/internal/builder"
	"flat-docs-service/internal/mapper"
	"flat-docs-service/pkg/db"
	flatbuffers "github.com/google/flatbuffers/go"
	"log"
)

const capacity = 100

type ReportService struct {
	ReportServiceInterface

	mongo  db.ReportClientRepository
	mapper mapper.FlatMapper
	bb     *builder.Pool
}

func NewReportService(mongo db.ReportClientRepository) *ReportService {
	return &ReportService{
		mongo:  mongo,
		mapper: *mapper.New(),
		bb:     builder.NewBuilderPool(capacity),
	}
}

func (sv *ReportService) Save(doc *sample.Document) error {
	parsedDoc := sv.mapper.MapToModel(doc)
	return sv.mongo.Save(*parsedDoc)
}

func (sv *ReportService) Find(limit int, offset int) (*[]byte, error) {
	docs, err := sv.mongo.Find(limit, offset)
	if err != nil {
		return nil, err
	}

	b := sv.bb.Get()
	defer sv.bb.Put(b)

	off := make([]flatbuffers.UOffsetT, len(docs))
	for i := 0; i < len(docs); i++ {
		parsedDoc := sv.mapper.MapToFlat(docs[i], b)
		off = append(off, parsedDoc)
	}

	sample.FindResponseStartDocsVector(b, len(off))
	for i := len(off) - 1; i >= 0; i-- {
		b.PrependUOffsetT(off[i])
	}
	docsVector := b.EndVector(len(off))

	sample.FindResponseStart(b)
	sample.FindResponseAddDocs(b, docsVector)
	response := sample.FindResponseEnd(b)
	b.Finish(response)
	fb := b.FinishedBytes()
	return &fb, nil
}

func (sv *ReportService) Validate(doc *sample.Document) error {
	code := string(doc.Department(new(sample.Department)).Code())
	log.Println("Department code:", code)
	if len(code) > 100 {
		return errors.New("ups, department code so big")
	}
	return nil
}

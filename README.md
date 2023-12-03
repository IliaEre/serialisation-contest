# serialisation-contest

The repository has three services:
1) json
2) proto (google api)
3) flutbuffer (google api)

Common stack:
1. golang
2. gin
3. mongodb 

Load testing:
1) Yandex tank + Yandex pandora for grpc and flatbuffer
2) locust (as option) for grpc
3) ghz (as option) for grpc

Tests and settings your can find 
```go
common/loadtest/
```

Serialisation tests:
- Json
```go
// BenchmarkCreateAndMarshal-10       168706       7045 ns/op
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
```
- proto
```go
// BenchmarkCreateAndMarshal-10       651063       1827 ns/op
func BenchmarkCreateAndMarshal(b *testing.B) {
 for i := 0; i < b.N; i++ {
  doc := CreateDoc()
  _ = doc.GetName()
  r, e := proto.Marshal(&doc)
  if e != nil {
   log.Fatal("problem with marshal")
  }

  nd := new(docs.Document)
  if proto.Unmarshal(r, nd) != nil {
   log.Fatal("problem with unmarshal")
  }
  _ = nd.GetName()
 }
}
```
- flatbuffer
```go
// BenchmarkCreateAndMarshalBuilderPool-10      1681384        711.2 ns/op
func BenchmarkCreateAndMarshalBuilderPool(b *testing.B) {
 builderPool := builder.NewBuilderPool(100)

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
```
Results:  
- j      168706       7045 ns/op  
- p      651063       1827 ns/op  
- f      1681384      711.2 ns/op

Flat is faster.


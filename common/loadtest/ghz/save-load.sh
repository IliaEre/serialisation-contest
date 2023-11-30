ghz --insecure -O html -o reports.html \
  --proto ../../../proto-docs-service/grpc/docs.proto \
  --call DocumentService.Save \
  -D report.json \
  -r 1000 -z 60s \
  localhost:84
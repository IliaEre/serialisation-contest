ghz --insecure -O html -o reports.html \
  --proto ../../../proto-docs-service/grpc/docs.proto \
  --call DocumentService.Save \
  -D report.json \
  --load-schedule=line --load-start=400 --load-end=1000 --load-step=10 -z 60s \
  localhost:84
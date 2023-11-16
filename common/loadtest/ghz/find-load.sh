ghz --insecure -O html -o reports_find.html \
  --proto ../../../proto-docs-service/grpc/docs.proto \
  --call DocumentService.GetAllByLimitAndOffset \
  -d '{"limit": 10, "offset":0}' \
  --load-schedule=line --load-start=400 --load-end=1000 --load-step=10 -z 60s \
  localhost:84
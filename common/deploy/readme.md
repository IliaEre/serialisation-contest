Grafana:
https://grafana.com/grafana/dashboards/3662-prometheus-2-0-overview/
https://grafana.com/grafana/dashboards/13240-go-metrics/



board:
rate(gin_requests_total{job!="prometheus"}[5m] ) * 60

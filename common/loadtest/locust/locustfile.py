import sys

from locust import HttpUser, task, between, tag


class ReportCase(HttpUser):
    wait_time = between(1, 2)
    host = "http://localhost"

    @task(1)
    @tag('save_report')
    def save(self):
        with open('save.bin', 'rb') as f:
            self.client.post('/report', data=f.read())

    @task(2)
    @tag('find_report')
    def find(self):
        with open('find.bin', 'rb') as f:
            self.client.post('/reports', data=f.read())

    def on_start(self):
        self.client.wait_time = between(0, 0)
        self.client.s = 400
        self.client.hatch_rate = (1000 - 400) / 60

        """on_start is called when a Locust start before any task is scheduled"""
        res = self.client.get("/metrics")
        if res.status_code != 200:
            print("Failed to find metrics on the server")
            sys.exit()

    def on_stop(self):
        self.client.wait_time = between(1, 2)
        self.client.hatch_rate = 1
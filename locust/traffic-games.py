from locust import HttpUser, task, between
from lector import Reader
import json


class Traffic(HttpUser):
    wait_time = between(0.1, 0.9)

    def on_start(self):
        self.reader = Reader()
        self.dataSend = {}
        self.reader.load()
        self.endpoint = "/match/addMatch"

    @task
    def initTraffic(self):
        self.dataSend = self.reader.getData()
        if self.dataSend is not None:
            json.dumps(self.dataSend)
            response = self.client.post(self.endpoint, json=self.dataSend)
            respuesta = json.loads(response.content)
            print(">> Reader-Locust-Response: ", respuesta)
        else:
            print(">> Reader-Locust: Envio de tráfico finalizado")
            self.stop(True)

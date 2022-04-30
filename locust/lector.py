from random import random, randrange
from sys import getsizeof
import json


class Reader:
    def __init__(self):
        self.array = []

    def getData(self):
        length = len(self.array)
        if length > 0:
            return self.array.pop(0)
        else:
            print(">> Reader-Locust: No hay más valores para leer en el archivo.")
            return None

    def load(self):
        print(">> Reader-Locust: Iniciando con la carga de datos")
        try:
            path = input("Reader-Locust: Ingrese la ruta del archivo JSON a cargar: ")
            with open(path, 'r') as dataFile:
                self.array = json.loads(dataFile.read())
            print(
                f'>> Reader-Locust: Datos cargados correctamente, {len(self.array)} datos -> {getsizeof(self.array)} bytes.')
        except Exception as e:
            print(f'>> Reader-Locust: No se cargaron los datos {e}')

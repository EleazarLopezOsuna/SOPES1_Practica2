# SOPES1_Practica2
![Arquitectura](https://drive.google.com/uc?export=view&id=1vJfSL6fjc3q9c_S6b-OVvEeNwDSFzGuo)
## Objetivos
- Experimentar y probar tecnologías Cloud Native que ayudan a desarrollar
sistemas distribuidos modernos.
- Diseñar estrategias de sistemas distribuidos para mejorar la respuesta
de alta concurrencia.
- Implementar contenedores y orquestadores en sistemas distribuidos.
- Medir la confiabilidad y el rendimiento en un sistema de alta
disponibilidad.
## Recursos utilizados
### Tecnologias
- [Google Cloud Platform](https://cloud.google.com/)
- [Docker](https://www.docker.com/)
- [Docker compose](https://docs.docker.com/compose/)
- [Kubernetes](https://kubernetes.io/)
- [Nodejs](https://nodejs.org/)
- [Kafka](https://kafka.apache.org/)
- [Strimzi](https://strimzi.io/)
- [gRPC](https://grpc.io/)
- [MongoDB](https://mongodb.com/)
- [Kubectl](https://kubernetes.io/docs/reference/kubectl/)
- [Locust](https://locust.io/)
- [MongoDB Compass](https://www.mongodb.com/products/compass)
### Lenguajes
- [Go](https://go.dev/)
- [Python](https://www.python.org/)
- [Javascript](https://developer.mozilla.org/en-US/docs/Web/JavaScript)
- [YAML](https://yaml.org/)
## Prerrequisitos
### Virtual Machine
Para poder hacer uso de ciertas tecnologias, necesitamos tener una VM ubicada en GCP la cual tiene esta configuracion inicial para poder hacer uso de Docker y Docker compose.
```bash
sudo apt-get update
sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
sudo usermod -aG docker $USER
newgrp docker

DOCKER_CONFIG=${DOCKER_CONFIG:-$HOME/.docker}
mkdir -p $DOCKER_CONFIG/cli-plugins
curl -SL https://github.com/docker/compose/releases/download/v2.4.1/docker-compose-linux-x86_64 -o $DOCKER_CONFIG/cli-plugins/docker-compose
chmod +x $DOCKER_CONFIG/cli-plugins/docker-compose
```
### Kubectl
Necesitamos Kubectl para poder interactuar con nuestro cluster.
```bash
wget curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.18.0/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
```
### Kubernetes Cluster
```bash
gcloud config set project <YOUR_PROYECT_NAME>
gcloud config set compute/zone us-central1-a
gcloud container clusters create k8s-demo --num-nodes=1 --tags=allin,allout --machine-type=n1-standard-2 --no-enable-network-policy
```
### Acceder desde nuestra VM al Cluster
```bash
gcloud container clusters get-credentials k8s-demo --zone=us-central1-c
```
### Instalacion de Strimzi
```bash
kubectl create namespace kafka
kubectl create -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka
```
## Archivos de Configuracion
### MongoDB
Para poder instalar MongoDB en nuestra VM haremos uso de docker compose. Para poder levantar MongoDB necesitamos ejecutar el siguiente comando `docker compose up -d` en la ubicacion donde tengamos nuestro archivo `docker-compose.yml`.
Los atributos entre <> son variables de entorno que pueden ser setteadas desde el YAML.
##### docker-compose.yml
```Docker
version: "3"
services:
  db:
    image: mongo
    container_name: dbmongo
    environment:
      MONGO_INITDB_ROOT_USERNAME: <YOUR_DB_USERNAME>
      MONGO_INITDB_ROOT_PASSWORD: <YOUR_DB_PASSWORD>
    ports:
      - "80:27017"
    restart: always
    volumes:
      - /home/($USER)/mongodata:/data/db
```
### API y gRPC Client
El archivo `api.yaml` permite levantar los siguientes elementos de Kubernetes.
- Contenedores
    - API
    - gRPC Client
- Servicios
    - LoadBalancer

Para poder levantar estos elementos debemos de ejecutar el siguiente comando `kubectl create -f api.yaml` en el directorio donde creemos el archivo.
Los atributos entre <> son variables de entorno que pueden ser setteadas desde el YAML.
##### api.yaml
```yaml
# API Deployment
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: api-server
  namespace: practica2-201700893
  name: api-server
spec:
  replicas: 1
  selector:
    matchLabels:
      app: api-server
  template:
    metadata:
      labels:
        app: api-server
    spec:
      containers:
        - image: jaredosuna/api_client_grpc_201700893
          name: api-server-client-grpc-201700893
          env:
            - name: SERVER_IP
              value: "<YOUR_gRPC_SERVER>"
            - name: SERVER_PORT
              value: "<YOUR_gRPC_PORT>"
            - name: API_PORT
              value: "<YOUR_API_PORT>"
          ports:
            - containerPort: <YOUR_API_PORT>
        - image: jaredosuna/api_server_grpc_201700893
          name: api-server-grpc-201700893
          ports:
            - containerPort: <YOUR_gRPC_PORT>
---
# API Service on port 3000
apiVersion: v1
kind: Service
metadata:
  name: api-server-svc
  namespace: practica2-201700893
  labels:
    app: api-server
spec:
  type: LoadBalancer
  ports:
    - port: <PORT_TO_EXPOSE>
      targetPort: <YOUR_API_PORT>
      protocol: TCP
  selector:
    app: api-server
```
### Kafka
El archivo `kafka.yaml` permite levantar los siguientes elementos de Kubernetes.
- Contenedores
    - zookeeper (x3)
    - kafka (x3)

Para poder levantar estos elementos debemos de ejecutar el siguiente comando `kubectl create -f kafka.yaml` en el directorio donde creemos el archivo.
Los atributos entre <> son variables de entorno que pueden ser setteadas desde el YAML.
##### kafka.yaml
```yaml
apiVersion: kafka.strimzi.io/v1beta2
kind: Kafka
metadata:
  name: my-cluster
  namespace: kafka
spec:
  kafka:
    version: 3.1.0
    replicas: 3
    listeners:
      - name: plain
        port: 9092
        type: internal
        tls: false
      - name: tls
        port: 9093
        type: internal
        tls: true
    config:
      offsets.topic.replication.factor: 3
      transaction.state.log.replication.factor: 3
      transaction.state.log.min.isr: 2
      log.message.format.version: "3.1"
      inter.broker.protocol.version: "3.1"
    storage:
      type: ephemeral
  zookeeper:
    replicas: 3
    storage:
      type: ephemeral
  entityOperator:
    topicOperator: {}
    userOperator: {}
```
### Consumer
El archivo `consumer.yaml` permite levantar los siguientes elementos de Kubernetes.
- Contenedores
    - consumer

Para poder levantar estos elementos debemos de ejecutar el siguiente comando `kubectl create -f consumer.yaml` en el directorio donde creemos el archivo.
Los atributos entre <> son variables de entorno que pueden ser setteadas desde el YAML.
##### consumer.yaml
```yaml
# Consumer Deployment
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: consumer
  namespace: practica2-201700893
  name: consumer
spec:
  replicas: 1
  selector:
    matchLabels:
      app: consumer
  template:
    metadata:
      labels:
        app: consumer
    spec:
      containers:
        - image: jaredosuna/api_consumer_grpc_201700893
          name: api-consumer-grpc-201700893
          env:
            - name: SERVER_PORT
              value: "<YOUR_gRPC_SERVER>"
            - name: KAFKA_HOST
              value: "<YOUR_KAFKA_HOST>"
            - name: KAFKA_PORT
              value: "<YOUR_KAFKA_PORT>"
            - name: KAFKA_TOPIC
              value: "<YOUR_KAFKA_TOPIC>"
            - name: MONGO_HOST
              value: "<YOUR_MONGO_HOST>"
            - name: MONGO_PORT
              value: "<YOUR_MONGO_PORT>"
            - name: MONGO_DB
              value: "<YOUR_MONGO_DATABASE>"
            - name: MONGO_COL
              value: "<YOUR_MONGO_COLLECTION>"
            - name: MONGO_USER
              value: "<YOUR_MONGO_USERNAME>"
            - name: MONGO_PASS
              value: "<YOUR_MONGO_PASSWORD>"
```
## Probando que nuestra arquitectura este funcionando de manera correcta
### Conectarnos a MongoDB
Para conectarnos a nuestra base de datos en MongoDB podemos hacer uso de MongoDB Compass utilizando el siguiente comando `mongodb://<YOUR_MONGO_USERNAME>:<YOUR_MONGO_PASSWORD>@YOUR_MONGO_HOST>:<YOUR_KAFKA_PORT>/?authSource=admin&readPreference=primary&directConnection=true&ssl=false` sustituyendo los valores dentro de <> con los respectivos.
### Utilizar Locust
Podemos hacer uso de Locust como load tester. Para poder hacer uso de esta herramienta debemos situarnos en el directorio correspondiente y ejecutar el siguiente comando.
```bash
locust -f traffic-games.py
```
Cuando nuestro archivo este ejecutado podremos acceder al siguiente sitio en nuestro navegador `http://localhost:8089/`. Seteamos los parametros y colocamos la `IP EXTERNA` de nuestro LoadBalancer. Para obtener esta direccion usamos el siguiente comando dentro de nuestra VM ` kubectl get services --namespace=practica2-201700893` y buscamos `EXTERNAL-IP` en donde `NAME = api-server-svc`.
Luego de llegar los campos que solicita Locust, debemos de regresar a la consola en donde ejecutamos Locust y colocar el nombre del archivo en donde tenemos los datos a cargar.
El archivo de carga debe de tener la siguiente estructura, sustituyendo los <>
```json
[
    {
        "game_id":<INTEGER>,
        "players":<INTEGER>
    },
    {
        "game_id":<INTEGER>,
        "players":<INTEGER>
    },
    ...
]
```
Luego de cargar el archivo podemos volver a MongoDB Compass y verificar los datos dentro de nuestra collection.
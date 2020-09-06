# go-amqp-worker

## Overview

a simple dockerised amqp consumer service template for golang


## Usage
### Running locally
```bash
git clone https://github.com/sanskarsharma/go-amqp-consumer.git
cd go-amqp-consumer
go run main.go
```
### Running via docker
```bash
git clone https://github.com/sanskarsharma/go-amqp-consumer.git
cd go-amqp-consumer
docker build -t go-amqp-consumer:v-local .
docker run -d go-amqp-consumer:v-local
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
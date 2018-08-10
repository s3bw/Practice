<h1 align="center">
    RabbitMQ Tutorial
</h1>

<h4 align="center">
    Covering the basics of creating messaging applications
</h4>

<p align="center">
    Source: https://www.rabbitmq.com/getstarted.html
</p>

## Install

```
pip install pika
```

## Setting up docker container

```bash
docker pull rabbitmq:latest
docker run -d -p 5672 --name rabbiting rabbitmq
```

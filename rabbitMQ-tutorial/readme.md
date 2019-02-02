<h1 align="center">
    RabbitMQ Tutorial
</h1>

<h4 align="center">
    Covering the basics of creating messaging applications
</h4>

<p align="center">
    Source: https://www.rabbitmq.com/getstarted.html
</p>

These tutorials are a great introduction to some messaging patterns.

- [Hello World](./tutorial_1/)
- [Round Robin](./tutorial_2/)
- [Pub/Sub](./tutorial_3/)
- [Routing](./tutorial_4/)
- [Topics](./tutorial_5/)
- [RPC](./tutorial_6/)

## Install

```
pip install pika
```

## Setting up docker container

```bash
docker pull rabbitmq:latest
docker run -d -p 5672 --name rabbiting rabbitmq
```

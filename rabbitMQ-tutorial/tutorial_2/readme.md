# Work Queues

## Round-robin dispatching

One of the advantages of using a Task Queue is the ability to easily
parallelise work. If we are building up a backlog of work, we can
add more workers and that way, scale easily.

## Usage

Start two instances of `worker.py`

```
python worker.py
```

Send messages to the workers

```
python new_task.py Message ....
```

## Message durability

We have learned how to make sure that even if the consumer dies,
the task isn't lost. But our tasks will still be lost if
RabbitMQ server stops.

When RabbitMQ quits or crashes it will forget the queues and
messages unless you tell it not to. Two things are required to
make sure that messages aren't lost: we need to mark both the
queue and messages as durable.

First, we need to make sure that RabbitMQ will never lose our
queue. In order to do so, we need to declare it as _durable_.

```python
channel.queue_declare(queue='hello', durable=True)
```

Although this command is correct by itself, it won't work in our
setup. That's because we've already defined a queue called
`hello` which is not durable. RabbitMQ doesn't allow you to
redefine an existing queue with different parameters and will
return an error to any program that tries to do that. But there
is a quick workaround - lets declare a queue with a different
name

```python
channel.queue_declare(queue='task_queue', durable=True)
```

This `queue_declare` change needs to be applied to both the
producer and consumer code.

At that point we're sure that the `task_queue` queue won't be
lost even if RabbitMQ restarts. Now we need to mark our
messages as persistent - by supplying a `delivery_mode` property
with a value `2`.

```python
channel.basic_publish(exchange='',
                      routing_key='task_queue',
                      body=message,
                      properties=pike.BasicProperties(
                        delivery_mode=2,  # make message persistent
                      ))
```

## Fair dispatch

You might noticed that the dispatching still doesn't work exactly
as we want. For example in a situation with two workers, when all
odd messages are heavy and even messages are light, one worker will
be constantly busy and the other one will do hardly any work. Well,
RabbitMQ doesn't know anything about that and will still dispatch
messages evenly.

This happens because RabbitMQ just dispatches a message when the
message enters the queue. It doesn't look at the number of
unacknowledged messages for a consumer. It just blindly dispatches
every n-th message to the n-th consumer.

In order to defeat that we can use the `basic.qos` method with the
`prefetch_count=1` setting. This tells RabbitMQ not to give more
than one message to a worker at a time. Or, in other words, don't
dispatch a new message to a worker until it has processed and
acknowledged the previous one. Instead, it will dispatch it to the
next worker that is not still busy.

```python
channel.basic_qos(prefetch_count=1)
```

#### Key Design Note - You'll want to monitor number of tasks in the queue.


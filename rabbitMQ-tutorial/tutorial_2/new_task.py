import sys

import pika


connection = pika.BlockingConnection(
    pika.ConnectionParameters('172.17.0.2', 5672)
)
channel = connection.channel()

channel.queue_declare(queue='task_queue', durable=True)


message = ' '.join(sys.argv[1:]) or "Hello World!"
channel.basic_publish(exchange='',
                      routing_key='task_queue',
                      body=message,
                      properties=pika.BasicProperties(
                        delivery_mode=2,  # make message persistent
                      ))


print(" [x] Sent %r" % message)
connection.close()

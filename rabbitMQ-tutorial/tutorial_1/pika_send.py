import pika

connection = pika.BlockingConnection(
    # To connect to a different machine we provide it's IP address
    pika.ConnectionParameters('172.17.0.2', 5672)
)
channel = connection.channel()

channel.queue_declare(queue='hello')

# In rabbitMQ a message can never be sent directly to the queue, it
# always need to go through an exchange. (see part 3)

channel.basic_publish(exchange='',
                      routing_key='hello',
                      body='Hello World!')
print(" [x] Sent 'Hello World!")

connection.close()

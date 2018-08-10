import pika

connection = pika.BlockingConnection(
    # To connect to a different machine we provide it's IP address
    pika.ConnectionParameters('172.17.0.2', 5672)
)
channel = connection.channel()

# It's good practice to repeat declaring the queue in both programs
channel.queue_declare(queue='hello')


def callback(ch, method, properties, body):
    """Receiving messages from the queue is more complex. It works by
        subscribing a 'callback' function to a queue. Whenever we
        receive a message, this 'callback' function is called by the
        Pika library. In our case this function will print on the
        screen the contents of the message.
    """
    print(" [x] Received %r" % body)

channel.basic_consume(callback,
                      queue='hello',
                      no_ack=True)

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()

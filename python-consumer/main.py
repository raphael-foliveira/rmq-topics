import os
import pika
from time import sleep


def message_callback(ch, method, properties, body):
    sleep(3)
    print(f"[{method.routing_key}] {body}")
    ch.basic_ack(delivery_tag=method.delivery_tag)


if __name__ == "__main__":
    amqp_host = os.environ.get("AMQP_HOST", "localhost")
    connection = pika.BlockingConnection(pika.ConnectionParameters(host=amqp_host))
    channel = connection.channel()

    channel.exchange_declare(exchange="topics", exchange_type="topic", durable=True)

    result = channel.queue_declare(queue="", durable=True, exclusive=True)
    queue_name = result.method.queue

    channel.queue_bind(exchange="topics", queue=queue_name, routing_key="topic.one")

    channel.basic_qos(prefetch_count=1)
    channel.basic_consume(
        queue=queue_name, on_message_callback=message_callback, auto_ack=False
    )
    channel.start_consuming()

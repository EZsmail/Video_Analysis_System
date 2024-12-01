import pika
import json

connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost', port=5672))
channel = connection.channel()

channel.queue_declare(queue='video_processing')

message = {
    "processing_id": "12345",
    "file_path": "/path/to/video.mp4"
}
channel.basic_publish(
    exchange='',
    routing_key='video_processing',
    body=json.dumps(message),
)

print("Sent task:", message)
connection.close()
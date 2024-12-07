import pika
import json
import time
from pymongo import MongoClient
import psycopg2
from psycopg2.extras import execute_values
import os


mongo_url = os.getenv("MONGO_URL", "mongodb://localhost:27017")
mongo_client = MongoClient(mongo_url)

mongo_db = mongo_client["video_processing_db"]
collection_results = mongo_db["collection_results"]



pg_connection = psycopg2.connect(
    dbname=os.getenv("POSTGRES_DB", "video_processing_db"),
    user=os.getenv("POSTGRES_USER", "postgres"),
    password=os.getenv("POSTGRES_PASSWORD", "password"),
    host=os.getenv("POSTGRES_HOST", "localhost"),
    port=os.getenv("POSTGRES_PORT", 5440)
)
pg_cursor = pg_connection.cursor()


def process_video(file_path):
    print(f"Processing video: {file_path}")
    time.sleep(5)
    return [["Part 1", "00:00", "00:30"], ["Part 2", "00:31", "01:00"]]

def update_status_in_postgresql(processing_id, status):
    query = "UPDATE processing_status SET status = %s WHERE processing_id = %s"
    pg_cursor.execute(query, (status, processing_id))
    pg_connection.commit()
    print(f"Updated status for {processing_id} to {status}")


def save_results_to_mongodb(processing_id, csv_data):
    document = {
        "_id": processing_id,
        "result_path": csv_data
    }
    collection_results.replace_one({"_id": processing_id}, document, upsert=True)
    print(f"Saved results to MongoDB for {processing_id}")


def on_message(channel, method, properties, body):
    message = json.loads(body)
    
    processing_id = message["processing_id"]
    file_path = message["file_path"]
    
    print(f"Received task: {processing_id} - {file_path}")
    
    try:
        update_status_in_postgresql(processing_id, "processing")
        

        csv_data = process_video(file_path)
        

        save_results_to_mongodb(processing_id, csv_data)
        

        update_status_in_postgresql(processing_id, "completed")
        
    except Exception as e:
        print(f"Error processing task {processing_id}: {e}")

        update_status_in_postgresql(processing_id, "failed")
    finally:
        channel.basic_ack(delivery_tag=method.delivery_tag)

rabbitmq_host = os.getenv("RABBIT_HOST", "localhost")
rabbitmq_port = os.getenv("RABBIT_PORT", 5672)
connection = pika.BlockingConnection(pika.ConnectionParameters(host=rabbitmq_host, port=rabbitmq_port))
channel = connection.channel()

channel.queue_declare(queue="video_processing")

channel.basic_consume(queue="video_processing", on_message_callback=on_message)
print("Waiting for messages...")
channel.start_consuming()
    
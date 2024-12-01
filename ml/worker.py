import pika 
import json
import time
import csv

def process_video(file_path):
    print(f"Processing video: {file_path}")
    time.sleep(5)
    return [["Part 1", "00:00", "00:30"], ["Part 2", "00:31", "01:00"]]

def save_csv(processing_id, csv_data):
    output_file = f"{processing_id}.csv"
    with open(output_file, mode="w", newline="") as file:
        writer = csv.writer(file)
        writer.writerow(["Part", "Start time", "End time"])
        writer.writerows(csv_data)
    print(f"Saved CSV: {output_file}")
    
def on_message(channel, method, properties, body):
    message = json.loads(body)
    
    processing_id = message["processing_id"]
    file_path = message["file_path"]
    
    print(f"Received task: {processing_id} - {file_path}")
    
    csv_data = process_video(file_path)
    
    save_csv(processing_id, csv_data)
    
    channel.basic_ack(delivery_tag=method.delivery_tag)
    
connection = pika.BlockingConnection(pika.ConnectionParameters(host="localhost", port="5672"))
channel = connection.channel()

channel.queue_declare(queue="video_processing")

channel.basic_consume(queue="video_processing", on_message_callback=on_message)
print("Waiting for messages...")
channel.start_consuming()

    
from celery import shared_task
import time
from app.services import ml_service


@shared_task
def process_video(file_path: str):
    
    
    # TODO: Add ML 
    time.sleep(10) #TODO: Remove
    
    video_id = "video_123"  
    
    # TODO: Change to db
    ml_service.save_results(video_id, {"csv_data": "example,data"})
    
    return video_id
from app.config.celery_config import celery_app
from app.ml import ML
from app.api.routes import video_data_repo, video_state_repo

# Инициализация клиента базы данных и репозиториев

@celery_app.task
def process_video(file_location: str, vid: str):
    
    # ML init
    ml_model = ML(vid, file_location)
    
    video_id, csv_data = ml_model.ml_template()
    
    # save csv to db
    video_data_repo.add_video_data(video_id, csv_data)
    
    # update state
    video_state_repo.update_video_state(video_id, "completed")
from fastapi import APIRouter, UploadFile, File, HTTPException
from fastapi.responses import StreamingResponse
from uuid import uuid4
from app.database import VideoStateRepository, VideoDataRepository
from app.database.mongo import MongoDBClient
from app.config import DATABASE_URL, UPLOAD_PATH
from pymongo import MongoClient
from pymongo.errors import ConnectionFailure
import pandas as pd
import io
import json
import os #dlya test ruchki



router = APIRouter()

# DB init
db_client = MongoDBClient(uri=DATABASE_URL, db_name="video_db")
video_state_repo = VideoStateRepository(db_client)
video_data_repo = VideoDataRepository(db_client)



@router.post("/upload-video/")
async def upload_video(file: UploadFile = File(...)):
    # Проверка подключения к БД
    try:
        client = MongoClient(DATABASE_URL)
        client.admin.command('ping')
        print("MongoDB доступен")
    except ConnectionFailure as e:
        print(f"Ошибка подключения к MongoDB: {e}")

    vid = str(uuid4())
    file_location = f"{UPLOAD_PATH}/{file.filename}"
    
    with open(file_location, "wb+") as file_object:
        file_object.write(file.file.read())

    # Добавление состояния видео (начальное состояние "processing")
    video_state_repo.add_video_state(vid, "processing")
    
    # TODO: Отправить задачу в Celery
    # process_video.delay(file_location, vid)
    
    return {"message": "Видео успешно загружено, началась обработка.", "vid": vid}



#get result by id
@router.get("/results/{vid}")
async def get_results(vid: str):
    try:
        result = video_data_repo.get_video_data(vid)
        return {"status": "completed", "data": result["csv"]}
    except ValueError:
        raise HTTPException(status_code=404, detail="Результаты обработки не найдены.")
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Ошибка получения данных: {str(e)}")


#get video status
@router.get("/video-status/{vid}")
async def get_video_status(vid: str):
    try:
        video_state = video_state_repo.get_video_state(vid)
        return {
            "video_id": vid,
            "status": video_state["state"]
        }
    except ValueError:
        raise HTTPException(status_code=404, detail="Видео не найдено.")
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Ошибка получения статуса видео: {str(e)}")



'''TODO
@router.post("/test_json_output/")
async def test_json_to_csv(file: UploadFile = File(...)):
    # Проверка на формат
    if not file.filename.endswith('.json'):
        raise HTTPException(status_code=400, detail="Файл должен быть JSON")

    contents = await file.read()
    
    try:
        # Загружаем JSON данные
        json_data = json.loads(contents)
        
        # Извлекаем первый словарь
        #first_dict = next(iter(json_data.values()))
        # Преобразуем JSON в DataFrame
        audio_df = pd.DataFrame(json_data['audio'])
        spekears_df = pd.DataFrame(json_data['spekears'])
        video_df = pd.DataFrame(json_data['video'])

        # Объединяем все DataFrame
        #combined_df = pd.concat([audio_df, spekears_df, video_df], axis=1)

        # Сохраняем в CSV
        audio_df.to_csv('output1.csv', index=False)
        spekears_df.to_csv('output2.csv', index=False)
        video_df.to_csv('output3.csv', index=False)
        temp_list = ['audio_df.csv', 'spekears_df.csv', 'video_df.csv']
        temp_dict = {"0": "", "1": "", "2": ""}

        
    except (ValueError, StopIteration) as e:
        raise HTTPException(status_code=400, detail="Неверный формат JSON или файл пуст")

    for i in range(3):
    # Преобразуем DataFrame в CSV
        csv_buffer = io.StringIO()
        data.to_csv(csv_buffer, index=False)
        csv_buffer.seek(0)

        # Сохраняем CSV файл на сервере
        csv_filename = temp_list[i]
        csv_filepath = os.path.join(UPLOAD_DIR, csv_filename)
        with open(csv_filepath, 'w') as f:
            f.write(csv_buffer.getvalue())

        # Возвращаем URL к загруженному CSV файлу
        file_url = f"http://localhost:8000/{UPLOAD_DIR}/{csv_filename}"
        temp_dict[str(i)] = file_url
        print(temp_dict)
    return temp_dict'''

@router.post("/test_json_output/")
async def test_json_to_csv(file: UploadFile = File(...)):
    try:
        # Можно добавить обработку данных, если нужно
        return file  # Возвращаем тот же JSON
        
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))
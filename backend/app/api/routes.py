from fastapi import APIRouter, UploadFile, File
from app.workers.task_queue import process_video
from app.services import ml_service
from uuid import uuid4

router = APIRouter()

# Load Video From User
@router.post("/upload-video/")
async def upload_video(file: UploadFile = File(...)):
    
    vid = uuid4()
    
    file_location = f"app/uploads/{file.filename}"
    with open(file_location, "wb+") as file_object:
        file_object.write(file.file.read())
    
    # TODO: Send task to Selery
    # process_video.delay(file_location, vid)
    
    return {"message": "Video uploaded successfully, processing started.", "vid": vid}


# Get Result Handler
@router.get("/results/{vid}")
async def get_results(vid: str):
    result = ml_service.get_results(vid)
    if result:
        return {"status": "completed", "data": result}
    else:
        return {"status": "processing"}


    
# Get Status Video
@router.get("/video-status/{vid}")
async def get_video_status(vid: str):

    video = videos_collection.find_one({"video_id": vid})
    
    if not video:
        return {"error": "Video not found"}
    
    return {
        "video_id": vid,
        "status": video.get("status"),
        "processing_results": video.get("processing_results", None)
    }
    
videos_collection = {}
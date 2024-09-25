from pydantic import BaseModel

class VideoModel(BaseModel):
    video_id: str
    file_path: str
    status: str
    result: dict = None
from app.database.mongo import MongoDBClient
from pymongo.errors import DuplicateKeyError, PyMongoError

class VideoStateRepository:
    def __init__(self, db_client: MongoDBClient, collection_name: str = "video_states"):
        self.collection = db_client.get_collection(collection_name)
    
    def add_video_state(self, video_id: str, state: str):
        document = {"_id": video_id, "state": state}
        try:
            self.collection.insert_one(document)
        except DuplicateKeyError:
            raise ValueError(f"Video state with ID '{video_id}' already exists.")
        except PyMongoError as e:
            raise RuntimeError(f"An error occurred while adding video state: {e}")
    
    def get_video_state(self, video_id: str) -> dict:
        video_state = self.collection.find_one({"_id": video_id})
        if not video_state:
            raise ValueError(f"No video state found for ID '{video_id}'")
        return video_state
    
    def update_video_state(self, video_id: str, new_state: str):
        result = self.collection.update_one({"_id": video_id}, {"$set": {"state": new_state}})
        if result.matched_count == 0:
            raise ValueError(f"No video state found for ID '{video_id}'")
    
    def delete_video_state(self, video_id: str):
        result = self.collection.delete_one({"_id": video_id})
        if result.deleted_count == 0:
            raise ValueError(f"No video state found for ID '{video_id}'")


class VideoDataRepository:
    def __init__(self, db_client: MongoDBClient, collection_name: str = "video_data"):
        self.collection = db_client.get_collection(collection_name)
    
    def add_video_data(self, video_id: str, csv_data: str):
        document = {"_id": video_id, "csv": csv_data}
        try:
            self.collection.insert_one(document)
        except DuplicateKeyError:
            raise ValueError(f"Video data with ID '{video_id}' already exists.")
        except PyMongoError as e:
            raise RuntimeError(f"An error occurred while adding video data: {e}")
    
    def get_video_data(self, video_id: str) -> dict:
        video_data = self.collection.find_one({"_id": video_id})
        if not video_data:
            raise ValueError(f"No video data found for ID '{video_id}'")
        return video_data
    
    def update_video_data(self, video_id: str, new_csv_data: str):
        result = self.collection.update_one({"_id": video_id}, {"$set": {"csv": new_csv_data}})
        if result.matched_count == 0:
            raise ValueError(f"No video data found for ID '{video_id}'")
    
    def delete_video_data(self, video_id: str):
        result = self.collection.delete_one({"_id": video_id})
        if result.deleted_count == 0:
            raise ValueError(f"No video data found for ID '{video_id}'")

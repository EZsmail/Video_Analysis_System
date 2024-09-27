from pymongo import MongoClient
from pymongo.errors import ConnectionFailure
from config import DATABASE_URL


try:
    client = MongoClient(DATABASE_URL)
    client.admin.command('ping')
    print("MongoDB доступен")
except ConnectionFailure as e:
    print(f"Ошибка подключения к MongoDB: {e}")
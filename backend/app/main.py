from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.api import routes
from app.database import MongoDBClient
from .config import DATABASE_URL


app = FastAPI()


# TODO: Change in prod
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"]
)

mongo_client = MongoDBClient(uri=DATABASE_URL, db_name="video_db")

app.include_router(routes.router)


@app.get("/")
def read_root():
    return {"message": "Backend is running!"}



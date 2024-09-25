from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.api import routes


app = FastAPI()


# TODO: Change in prod
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


# Подключаем маршруты
app.include_router(routes.router)


@app.get("/")
def read_root():
    return {"message": "Backend is running!"}

from celery import Celery
from app.config.config import CELERY_URL
import os


celery = Celery(
    "ml",
    broker=CELERY_URL
)


# Load tasks
celery.autodiscover_tasks(["app.workers.task_queue"])

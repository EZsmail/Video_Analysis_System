from celery import Celery

celery_app = Celery(
    'app',
    broker='amqp://guest:guest@localhost:5672//',  # Подключение к RabbitMQ через localhost
    backend='redis://localhost:6379/0',  # Подключение к Redis через localhost
)

# Конфигурация задач Celery
celery_app.conf.update(
    task_routes={
        'app.tasks.*': {'queue': 'default'},
    },
    result_backend='redis://localhost:6379/0',
    task_serializer='json',
    accept_content=['json'],
    result_serializer='json',
    timezone='UTC',
    enable_utc=True,
)

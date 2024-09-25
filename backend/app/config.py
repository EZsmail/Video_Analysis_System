from envparse import Env

env = Env()

DATABASE_URL = env.str(
    'MONGOURL',
    default="postgresql+asyncpg://admin:root@0.0.0.0:5438/postgres"
)

CELERY_URL = env.str(
    "CELERY_URL",
    default="pyamqp://guest:guest@localhost//"
)

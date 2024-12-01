from envparse import Env

env = Env()

DATABASE_URL = env.str(
    'DB_URL',
    default="mongodb://qioto:qwerty@localhost:27017/"
)

CELERY_URL = env.str(
    "CELERY_URL",
    default="pyamqp://guest:guest@localhost//"
)

UPLOAD_PATH = env.str(
    "UPLOAD_PATH",
    default="../data/upload/"
)
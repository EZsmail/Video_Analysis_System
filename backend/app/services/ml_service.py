import time


# save res to mongo
def save_results(video_id: str, data: dict): 
    # TODO: add save to db
    results_db[video_id] = data


#get data
def get_results(video_id: str):
    return results_db.get(video_id)

# init db
results_db = {}
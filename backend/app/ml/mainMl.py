import time
import csv
import io

# TODO: Change to ML
def ml_template(_id: str, url: str) -> tuple:
    data = [
        ['ID', 'Name', 'Age'],
        [1, 'John', 28],
        [2, 'Jane', 22]
    ]
    

    output = io.StringIO()
    csv_writer = csv.writer(output)
    

    csv_writer.writerows(data)
    

    csv_content = output.getvalue()
    

    output.close()
    
    time.sleep(15)
    
    return _id, csv_content


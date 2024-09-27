import time
import csv
import io

class ML:
    def __init__(self, _id, url):
        self._id = _id
        self.url = url
        
    # TODO: Change to ML
    def ml_template(self) -> tuple:
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
        
        return self._id, csv_content


# 🎥 Video Processing System

This project is a video processing system that enables users to upload videos, processes them into meaningful segments using an ML service, and displays the results via a web interface.

## 🚀 Key Features

-	Frontend:
  -	A simple and user-friendly web interface for uploading videos and viewing results.
-	Backend:
  -	Built with Go for performance and scalability.
  -	Powered by Gin for routing and Zap for structured logging.
-	Asynchronous Processing:
  -	Video tasks are queued using RabbitMQ, ensuring smooth processing.
  -	Background worker processes tasks with a Python ML service.
-	Data Storage:
  -	MongoDB for storing processed video results.
  -	PostgreSQL for tracking video processing statuses.

## 📦 Tech Stack

| Component | Technology |
| --- | --- |
| Frontend | HTML, CSS, JS |
| Backend | Go (Gin, Zap) |
| Worker | Python (ML Models) |
| Message Queue | RabbitMQ |
| Databases | MongoDB, PostgreSQL |
| Cache | Redis |

## 🖥️ Usage

🛠️ How to Use This Project

1.	Clone the Repository
```
git clone https://github.com/your-username/your-repo-name.git  
cd your-repo-name  
```  
2.	Install Docker  
Ensure Docker and Docker Compose are installed. Get them from Docker’s official site.  
3.	Check Ports  
Ensure these ports are free or update docker-compose.yml if needed:  
	•	Frontend: 8000  
	•	Backend: 8080  
	•	RabbitMQ: 5672 (15672 for management)  
	•	MongoDB: 27017  
	•	PostgreSQL: 5432  
4.	Start Services  
Run:  
```
docker-compose up --build  
```  
This starts all containers and links services.  
5.	Access the App  
-	Open http://localhost:8000 (or your specified frontend port), upload a video, and wait for processing.  
6.	View Results  
-	Use the Results section to input the Processing ID and see the segmented data.  
7.	Stop Services  
Run:  
```
docker-compose down  
```

## 📈 Future Enhancements

-	Full integration with advanced ML models for better video analysis.  
-	GPU acceleration to improve processing speed.  
-	Enhanced frontend visualization for a richer user experience.  

## 🤝 How to Contribute

We welcome contributions, suggestions, and feature requests! Check out the issues page for ideas.  
  
1.	Fork this repository.  
2.	Create a new branch for your feature or fix.  
3.	Make your changes and commit them.  
4.	Push your branch to your forked repository.  
5.	Submit a pull request for review.  
 

# 🎥 Video Analysis System

This project is a video analysis system that allows users to upload videos, break them into semantic parts, highlight the essence, find non-residents using an ML service and display the results via a web interface.

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
| Backend | Go (Gin, Zap) |
| Worker | Python (Tensorflow, Whisper, GPT2, Torch) |
| Frontend | HTML, CSS, JS |
| Message Queue | RabbitMQ |
| Databases | MongoDB, PostgreSQL |
| Cache | Redis |

## 🖥️ Installation
  
```bash
git clone https://github.com/EZsmail/ww.git
cd ww
docker compose up --build
```
  
## 🖥️ Usage
  
Go to the address `http://localhost:8000`
The port may change depending on which one you specified in docker-compose.yaml.
  
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
 

# Dockerfile for Python service
FROM python:3.11

WORKDIR /app

COPY . .

RUN pip install Flask pymongo

CMD ["python", "app.py"]

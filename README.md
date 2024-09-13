# Restaurant App

This project is a restaurant application that includes a Go server for handling user and item management and a Python service for processing orders. The application uses MongoDB as the database.

## Table of Contents

- [Technologies](#technologies)
- [Setup](#setup)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Technologies

- Go
- Python
- MongoDB
- Gin (Go Web Framework)
- Flask (Python Web Framework)
- Docker

## Setup

### Prerequisites

- Go 1.16+
- Python 3.8+
- Docker
- MongoDB

### Environment Variables

Create a `.env` file in the root directory and add the following:

```env
MONGO=<your_mongo_connection_string>
JWT_SECRET = <Your Secret Key>
```

# Go Server
### Navigate to the go_server directory:  
```bash
cd go_server
```

### Install dependencies:
```bash 
go mod tidy
```

### Run the server:
```bash
go run main.go
```

# Python Order Service

### Navigate to the pythonOrderService directory:
```bash
cd pythonOrderService
```

### Create a virtual environment and activate it:
```bash
python -m venv venv
source venv/bin/activate  # On Windows use `venv\Scripts\activate`
```

### Install dependencies:
```bash
pip install -r requirements.txt
```

### Run the service:
```bash
python app.py
```




# go-kafka-event-stream

A GoLang-based event-driven microservices application demonstrating real-time data streaming with Apache Kafka.

## Overview
This project implements a scalable producer-consumer system:
- **Producer**: A Go server using Fiber to expose a REST API (`/api/v1/comments`) that accepts JSON payloads and publishes them to a Kafka topic (`comments`).
- **Consumer**: A Go server that subscribes to the `comments` topic and logs messages in real-time.
- **Purpose**: Showcases asynchronous communication in microservices using Kafka for high-throughput, decoupled data streaming.
- **Tech Stack**: GoLang, Apache Kafka, Zookeeper, Docker, Fiber, Sarama, godotenv.

## Features
- REST API for posting comments to Kafka.
- Real-time message consumption and logging.
- Containerized Kafka and Zookeeper setup with Docker Compose.
- Environment-based configuration using `.env`.

## Prerequisites
- Go (1.18 or higher)
- Docker and Docker Compose
- Git
- Curl (for testing)

## Setup Instructions
1. Clone the repository:
   ```bash
   git clone https://github.com/salamafazlul/go-kafka-event-stream.git
   cd go-kafka-event-stream
   ```
2. Install Go dependencies:
   ```bash
   go get github.com/gofiber/fiber/v2
   go get github.com/IBM/sarama
   go get github.com/joho/godotenv
   go mod tidy
   ```
3. Create a `.env` file:
   ```bash
   echo "KAFKA_BROKERS=localhost:9092" > .env
   ```
4. Start Kafka and Zookeeper:
   ```bash
   docker-compose up -d
   ```
5. Run the producer:
   ```bash
   cd producer
   go run producer.go
   ```
6. Run the consumer (in a new terminal):
   ```bash
   cd worker
   go run worker.go
   ```
7. Test the application:
   ```bash
   curl -X POST http://localhost:3000/api/v1/comments -H "Content-Type: application/json" -d '{"text":"message one"}'
   ```

## Expected Output
- **Producer**: `{"success":true,"message":"Comment pushed successfully","comment":{"text":"message one"}}`
- **Consumer**: `Received message Count: 1: | Topic comments) | Message({"text":"message one"})`

## Project Structure
```
go-kafka-event-stream
├── producer
│   └── producer.go      # REST API server to publish to Kafka
├── worker
│   └── worker.go        # Consumer to process Kafka messages
├── docker-compose.yml   # Kafka and Zookeeper configuration
├── .env                 # Environment variables (not committed)
├── .gitignore           # Excludes unnecessary files
├── go.mod               # Go module dependencies
├── go.sum               # Dependency checksums
├── README.md            # Project documentation
```

## Notes
- Uses standard Kafka port `9092` for compatibility.
- Configurable via `.env` for flexible deployment.
- Tested in WSL2; for stability, consider a VirtualBox Ubuntu VM if Docker issues arise.
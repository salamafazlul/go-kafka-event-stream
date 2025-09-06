# go-kafka-event-stream
A GoLang-based event-driven microservices application using Apache Kafka for real-time data streaming

## Overview
- **Producer**: Go server with Fiber exposing a REST API to send comments to a Kafka topic.
- **Consumer**: Go server consuming and logging messages from the Kafka topic.
- **Tech Stack**: GoLang, Kafka, Zookeeper, Docker, Fiber, Sarama.

## Setup
1. Install Go, Docker, Docker Compose, Git, Curl.
2. Run `docker-compose up -d` to start Kafka/Zookeeper.
3. Run `go run producer/producer.go` and `go run worker/worker.go`.
4. Test: `curl -X POST http://localhost:3000/api/v1/comments -H "Content-Type: application/json" -d '{"text":"message one"}'`

## Prerequisites
- Go 1.18+
- Docker & Docker Compose
- Git, Curl
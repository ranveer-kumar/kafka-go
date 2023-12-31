## Kafka Go Project

This project guide for working with Apache Kafka, specifically with AWS Managed Streaming for Apache Kafka (AWS MSK), within a Go application. It effectively demonstrates the utilization of the Kafka franz-go library to perform essential tasks such as producing and consuming messages from Kafka topics. By exploring this project, you will gain an understanding of how to seamlessly integrate Kafka messaging capabilities into your Go-based projects.

### Project Structure


#### Directories and Files

- `cmd/`: Contains the main entry point of the application for producer (cmd/producer/main.go) and for consumer (cmd/consumer/main.go).
- `config/`:  Configuration-related code.
- `pkg/kafka`:  Kafka clinet-related code.
- `producer/`: Contains the Kafka producer code.
  - `producer.go`: Kafka producer logic.
  - `handlers.go`: HTTP request handlers for the producer component.
- `consumer/`: Contains the Kafka consumer code.
  - `consumer.go`: Kafka consumer logic.
  - `handlers.go`: HTTP request handlers for the consumer component.
- `go.mod` and `go.sum`: Go module files.



### Getting Started

#### Prerequisites

Before running the project, make sure you have the following installed:

  

Go (https://golang.org/)

Kafka broker (https://kafka.apache.org/) or AWS MSK should be configured with SASL_SSL and open publicly


### Installation

1.  Clone this repository:
    
```shell
git clone https://github.com/yourusername/kafka-go.git
cd kafka-go
```

2.  Install dependencies:
    
```
go mod tidy
```
    

## Usage

### Running the Kafka Consumer

To run the Kafka consumer, execute the following command:

```
go run cmd/consumer/main.go
```

This will start the consumer, which will continuously consume messages on console from the Kafka topic.

### Running the Kafka Producer

To run the Kafka producer, execute the following command:

```
go run cmd/producer/main.go
```

This will start an HTTP server that listens for incoming requests to send messages to a Kafka topic.

### Sending Messages

To send a message to Kafka, make an HTTP GET request to **http://localhost:8080/send** with the following query parameters:

-   **message**: The message you want to send.
-   **topic**: The Kafka topic to which you want to send the message.

Example:
```curl "http://localhost:8080/send?message=Hello%20Kafka&topic=mytopic"```

## Contributing

Contributions to this project are welcome. Feel free to open issues and pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/ranveer-kumar/kafka-go.git) file for details.


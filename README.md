# Football Data

This project is a Go application that fetches and processes football data.

## Prerequisites

Ensure you have Go installed on your machine. This project uses Go Modules for dependency management.

## Environment Variables

The application requires the following environment variables:

- `RAPID_API_KEY`: Your RapidAPI key.
- `ENVIRONMENT`: The environment in which the application is running. For local development, this should be `local`.
- `KAFKA_USERNAME`: Your Kafka username.
- `KAFKA_PASSWORD`: Your Kafka password.
- `KAFKA_HOST`: The host address of your Kafka instance.
- `KAFKA_TOPIC_FIXTURE`: The Kafka topic for fixtures.
- `KAFKA_TOPIC_TEAMS`: The Kafka topic for teams.
- `KAFKA_TOPIC_STANDING`: The Kafka topic for standings.

You can set these environment variables in a `.env` file in the root of the project. Use the provided `.env.dist` file as a template.

## Build

To build the application, run the following command:

```
make football_data
```

## Run

To run the application, first ensure that all environment variables are set, then execute the binary:

```
./footballDataApp
```

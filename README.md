# go-expert-stress-test

This is a command-line interface (CLI) tool written in Go for performing load tests on web services.
It allows you to simulate a high volume of traffic to a specified API, helping you measure and evaluate its performance and resilience under stress. The tool provides a detailed and easy-to-read report upon completion.

---

## Table of Contents
- [Project Structure](#project-structure)
- [How It Works](#how-it-works)
- [Getting Started](#getting-started)
- [Running the Project](#running-the-project)

---

## Project Structure

```text
go-expert-stress-test/
|-- cmd/
|   |-- server/                 # Main entry point for the HTTP server
|       |-- [main.go]
|-- internal/
|   |-- stresstest/             # Report logic
|   |   |-- [report.go]
|   |   |-- [worker.go]         # Worker logic
|-- [Dockerfile]                # Docker configuration to build the image
|-- [go.mod]
|-- [README.md]
```

## How It Works
1.  **Parameter Parsing**: The application parses command-line flags (`--url`, `--requests`, `--concurrency`) to configure the test.
2.  **Load Distribution**: The total number of requests is distributed among multiple Go routines, which act as workers, based on the specified concurrency level.
3.  **Concurrent Execution**: Each worker concurrently sends HTTP requests to the target URL, measuring the latency and collecting the HTTP status code for each response.
4.  **Report Generation**: Once all requests are completed, the application processes the collected data to generate a comprehensive report, including total time, status code distribution, and average latency.

## Getting Started
Prerequisites
* Go 1.18+
* Docker

Clone the repository
```bash
git clone https://github.com/ricardolindner/go-expert-stress-test.git
cd go-expert-stress-test
```

Download the dependencies:
```bash
go mod tidy
```

## Running the Project

The CLI tool expects the following parameters to run the load test:

|Flag|Type|Description|Example|
|:---|:---|:---|:---|
|`--url` |`string`|The target URL to test.|`--url=https://google.com`|
|`--requests`|`int`|The total number of requests.|`--requests=1000` |
|`--concurrency`|`int`| The number of simultaneous workers.|`--concurrency=10`|

### 1. Running with the Go CLI

From the project's root directory, execute the following command:
```bash
go run ./cmd/server/main.go --url=http://google.com --requests=1000 --concurrency=10
```

### 2. Running with docker
Build the Docker image:
```bash
docker build -t stress-test-cli .
```
Run the load test:
```bash
docker run --rm stress-test-cli --url=http://google.com --requests=1000 --concurrency=10
```

## Report Sample

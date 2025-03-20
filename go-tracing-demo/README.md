# Distributed Tracing with DataDog in Go

This example demonstrates how to implement **distributed tracing** in a Go application using **DataDog**. It consists of two services:
- **Service A**: An HTTP server that calls Service B.
- **Service B**: An HTTP server that simulates some work.

The traces from both services are sent to a **DataDog Agent**, which forwards them to the DataDog backend for visualization.

---

## Prerequisites

1. **Go**: Install Go from [https://golang.org/dl/](https://golang.org/dl/).
2. **Docker**: Install Docker from [https://www.docker.com/get-started](https://www.docker.com/get-started).
3. **DataDog Account**: Sign up at [https://www.datadoghq.com/](https://www.datadoghq.com/) and get your **API Key**.

---

## Setup

### 1. Clone the Repository
Clone this repository to your local machine:
```bash
git clone <repository-url>
cd go-tracing-demo
```

### 2. Install Dependencies
Navigate to each service directory and initialize the Go modules:
```bash
cd service-a
go mod tidy

cd ../service-b
go mod tidy
```

### 3. Run the DataDog Agent
Start the DataDog Agent using Docker:
```bash
docker run -d --name datadog-agent \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  -v /proc/:/host/proc/:ro \
  -v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro \
  -p 127.0.0.1:8126:8126/tcp \
  -e DD_API_KEY=<YOUR_DATADOG_API_KEY> \
  -e DD_SITE="<DD_SITE_LINK>" \
  datadog/agent:latest
```
Replace `<YOUR_DATADOG_API_KEY>` with your actual DataDog API key.

---

## Running the Services

### 1. Start Service B
Navigate to the `service-b` directory and run the service:
```bash
cd service-b
go run main.go
```
Service B will start listening on `http://localhost:8081`.

### 2. Start Service A
Navigate to the `service-a` directory and run the service:
```bash
cd service-a
go run main.go
```
Service A will start listening on `http://localhost:8080`.

---

## Testing the Setup

### 1. Make a Request to Service A
Use `curl` or a browser to make a request to Service A:
```bash
curl http://localhost:8080/
```
You should see the response:
```
Hello from service-a!
```

### 2. Check Logs
- Service A will log the response from Service B.
- Service B will log the incoming request.

---

## Visualizing Traces in DataDog

1. Log in to your DataDog account.
2. Go to the **APM > Traces** section.
3. You should see a **single trace** that includes spans from both Service A and Service B.
4. Use the **Service Map** to visualize the relationship between the services.
5. Explore the **Trace View** to see the timing and details of each span.

---

## Project Structure

```
go-tracing-demo/
├── service-a/
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── service-b/
│   ├── main.go
│   ├── go.mod
│   └── go.sum
└── README.md
```

---

## Key Features

- **Distributed Tracing**: Traces are propagated from Service A to Service B.
- **DataDog Integration**: Traces are sent to the DataDog backend for visualization.
- **HTTP Instrumentation**: Automatic trace context propagation using DataDog's HTTP client and server instrumentation.

---

## Troubleshooting

### 1. DataDog Agent Not Running
Ensure the DataDog Agent container is running:
```bash
docker ps
```
If the container is not running, check the logs:
```bash
docker logs datadog-agent
```

### 2. Traces Not Appearing in DataDog
- Verify that the DataDog API key is correct.
- Ensure the DataDog Agent is configured to allow non-local traffic (`DD_APM_NON_LOCAL_TRAFFIC=true`).
- Check the logs of Service A and Service B for errors.

---

## Credits

This project was developed with the assistance of **DeepSeek**, an AI-powered coding assistant. Special thanks to DeepSeek for providing guidance, code examples, and troubleshooting support.

---

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
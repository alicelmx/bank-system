# Build Stage
# Use an official Go runtime as a parent image
FROM golang:1.21.1-alpine3.18 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project's source code into the container
COPY . .

# Build the Go application
RUN go build -o main main.go

# Run Stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main/ .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

# Expose a port for the application to listen on (change as needed)
EXPOSE 9998

# Define the command to run your application
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]

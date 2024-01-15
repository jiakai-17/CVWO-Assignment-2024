FROM golang:alpine3.19 AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Install packages
RUN go mod download

# Compile the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-backend

FROM alpine:3.19

COPY --from=builder /docker-backend /

# Expose port 9090
EXPOSE 9090

# Run the binary
CMD ["/docker-backend"]

FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o martian-robots main.go

# Run the application
CMD ["./martian-robots", "input.txt"]

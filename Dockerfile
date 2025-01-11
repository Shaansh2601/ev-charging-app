# Use the specific Go version
FROM golang:1.23.3-alpine

# Install necessary tools
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

COPY client ./client

# Copy only the module files first (to cache dependencies)
COPY backend/go.mod backend/go.sum ./backend/
WORKDIR /app/backend
RUN go mod download

# Copy the entire backend directory into the container
COPY backend ./  

# Build the Go binary
RUN go build -o EV-app ./cmd/main.go

# Expose the port your app runs on
EXPOSE 8000

# Command to run the application
CMD ["./EV-app"]



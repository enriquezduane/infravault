FROM golang:alpine

WORKDIR /usr/src/app

# Install wgo for hot reloading (watches file changes and restarts)
RUN go install github.com/bokwoon95/wgo@latest

# Copy go.mod first for better caching
COPY go.mod ./

# Download dependencies (this step is cached if go.mod doesn't change)
RUN go mod download

# Copy source code
COPY . .

# Use wgo for development hot reloading
CMD ["wgo", "run", "main.go"]

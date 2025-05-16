# syntax=docker/dockerfile:1

### Build stage
FROM golang:1.23.0-alpine AS builder

# Install git for go modules if needed
RUN apk add --no-cache git

# Set destination for COPY
WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .


# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/main.go

### Final stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Install Terraform dependencies
RUN apk add --no-cache wget unzip curl

# Install Terraform
ENV TERRAFORM_VERSION=1.11.4
RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    mv terraform /usr/local/bin/ && \
    rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip

# Copy binary from builder
COPY --from=builder /app/server .

#COPY --from=builder /app/config.json .

COPY --from=builder /app/step-templates /app/step-templates

# Set up plugins directory
RUN mkdir -p /app/plugins && chmod 777 /app/plugins

# Expose the port
EXPOSE 5000


# Run the server
CMD ["./server"]

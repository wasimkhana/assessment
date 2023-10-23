# Makefile for building Docker image, creating network, running the container, and running the Go app

.PHONY: all build-image create-network run-container install-dependencies run-app clean

# Define variables
DOCKER_IMAGE_NAME = my-docker-image
DOCKER_NETWORK_NAME = mynetwork
CONTAINER_NAME = my-container
GO_APP_NAME = main.go

all: build-image create-network run-container

# Build Docker image from Dockerfile
build-image:
	docker build -t $(DOCKER_IMAGE_NAME) .

# Create a custom Docker network with a specific subnet
create-network:
	docker network create --subnet=172.18.0.0/16 $(DOCKER_NETWORK_NAME)

# Run the container and execute the Go app within the container
run-container:
	docker run --name $(CONTAINER_NAME) --network $(DOCKER_NETWORK_NAME) --ip 172.18.0.2 $(DOCKER_IMAGE_NAME) go run $(GO_APP_NAME)

# Install Go dependencies (if needed)
install-dependencies:
	go get

# Clean up: Stop and remove the container, remove the network, and remove the Docker image
clean:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true
	docker network rm $(DOCKER_NETWORK_NAME) || true
	docker rmi $(DOCKER_IMAGE_NAME) || true

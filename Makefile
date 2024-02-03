build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go

# Define variables
IMAGE_NAME := backend-payment-service
DOCKER_REPO := stanthikun802
DOCKER_TAG := latest

# Build the Docker image
docker-build:
	docker build -t $(IMAGE_NAME) .

stop:
	docker stop backend-payment-service && docker rm backend-payment-service

drun:
	docker run -d --name backend-payment-service -p 3101:8080 backend-payment-service

# Tag the Docker image
tag:
	docker tag $(IMAGE_NAME) $(DOCKER_REPO)/$(IMAGE_NAME):$(DOCKER_TAG)

# Push the Docker image to the repository
push:
	docker push $(DOCKER_REPO)/$(IMAGE_NAME):$(DOCKER_TAG)

# Remove the locally built Docker image
clean:
	docker rmi $(IMAGE_NAME)

# Chain tasks together
all: docker-build tag push clean

local: docker-build stop run

# Define phony targets to avoid conflicts with file names
.PHONY: docker-build tag push clean all stop run local drun
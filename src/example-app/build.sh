#!/usr/bin/env bash

VERSION=$(cat version)
IMAGE_NAME="abhijitwakchaure/example-app"

echo -e "\nDeleting old image..."
docker rmi -f ${IMAGE_NAME}:latest
echo -e "\nDeleting old image...done\n"

echo -e "\nBuilding go code..."
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app-server .
echo -e "\nBuilding go code...done\n"

echo -e "\nBuilding docker image..."
docker build -t ${IMAGE_NAME}:latest .
echo -e "\nBuilding docker image...done\n"

echo -e "\nTagging docker image with v${VERSION}..."
docker tag ${IMAGE_NAME}:latest ${IMAGE_NAME}:${VERSION}
echo -e "\nTagging docker image with v${VERSION}...done\n"

echo -e "\nPushing docker image..."
docker push ${IMAGE_NAME}:latest
docker push ${IMAGE_NAME}:${VERSION}
echo -e "\nPushing docker image...done\n"

echo -e "\nCleaning up..."
rm -f app-server
echo -e "\nCleaning up...done\n"

echo -e "\nStarting docker container for ${IMAGE_NAME}:${VERSION}..."
docker run --rm -it -p 8080:8080 ${IMAGE_NAME}:${VERSION}
echo -e "\nStarting docker container for ${IMAGE_NAME}:${VERSION}...done\n"
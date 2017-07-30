#! /bin/bash

# Pull the most recent version
git pull

# Overwrite the dockerfile
cp Dockerfile.prod Dockerfile

# Overwrite the docker compose file
cp docker-compose.yml.prod docker-compose.yml

# Take our server down
docker-compose down

# Build our new up
docker-compose build

# Finally deploy it
docker-compose up -d

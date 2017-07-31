#! /bin/bash

# This script is designed to be used in production
# it takes down current prod, cleans our old containers
# rebuilds the app and then brings the api back up

# Reset everything so that git will shut up
git reset --hard origin/master

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

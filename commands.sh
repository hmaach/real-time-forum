#!/bin/bash

# Build a new image
docker build --no-cache -f dockerfile -t forum-img .

# Run a new container
docker run -d -p 8080:8080 --name forum-con forum-img

docker logs forum-con

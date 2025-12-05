#!/bin/bash

# storage server
docker build -t eugenebalykov/health-storage:latest -f cmd/server/Dockerfile .
docker push eugenebalykov/health-storage:latest

# storage client
docker build -t eugenebalykov/health-storage-client:latest -f cmd/client/Dockerfile .
docker push eugenebalykov/health-storage-client:latest

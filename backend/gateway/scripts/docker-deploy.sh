#!/bin/bash

# storage client
docker build -t eugenebalykov/gateway:latest -f Dockerfile .
docker push eugenebalykov/gateway:latest

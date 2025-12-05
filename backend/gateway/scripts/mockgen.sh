#!/bin/bash

# user service client mock
mockgen -source=./internal/services/users/client.go \
        -destination=./internal/mocks/mock_user_service.go \
        -package=mocks
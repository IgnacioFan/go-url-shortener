#!/bin/bash

ENV_CONTENT=$(cat <<EOF
# DB
POSTGRES_HOST=go-url-shortener-db
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=go-url-shortener
POSTGRES_PORT=5432

# Redis
REDIS_HOST=go-url-shortener-redis
REDIS_PORT=6379

EOF
)
echo "$ENV_CONTENT" > .env

echo "Env variables and .env file generated successfully!"
echo "Now, you can edit the .env file to fit your needs!"

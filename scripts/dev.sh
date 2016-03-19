#!/usr/bin/env bash

docker rm -fv price-history
docker rm -fv ph-db

docker build -f docker/Dockerfile -t price-history . && \

docker run --name ph-db -e POSTGRES_PASSWORD=pass1234 -e POSTGRES_DB=postgres -e POSTGRES_USER=postgres -d postgres
sleep 10s #need to wait while postgresql starting up
docker run --name price-history -p 8080:8080 -it -v $(pwd):/go/src/app --link ph-db:postgres price-history bash -c "./scripts/migrate.sh && /bin/bash"


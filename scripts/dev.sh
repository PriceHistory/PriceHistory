#!/usr/bin/env bash

docker rm -fv price-history
docker rm -fv ph-db

docker build -f docker/Dockerfile -t price-history . && \

docker run --name ph-db -e POSTGRES_PASSWORD=pass1234 -e POSTGRES_DB=postgres -e POSTGRES_USER=postgres -d postgres
docker run --name price-history -p 127.0.0.1:8080:8080 -it -v $(pwd):/go/src/app --link ph-db:postgres price-history /bin/bash
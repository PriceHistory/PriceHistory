#!/usr/bin/env bash

docker rm -fv dev-price-history

docker build -f docker/Dockerfile -t price-history . && \

docker run --name price-history -it -v $(pwd):/go/src/app price-history /bin/bash
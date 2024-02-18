#!/usr/bin/env bash

echo "stop metrics servers"
docker stop grafana prometheus

echo "remove metrics servers"
docker rm grafana prometheus
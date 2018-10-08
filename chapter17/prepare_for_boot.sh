#!/bin/bash

docker pull alpine
docker run --name alpine alpine
docker export alpine > alpine.tar
docker rm alpine
mkdir -p rootfs
tar -C rootfs -xvf alpine.tar

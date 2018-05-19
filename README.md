# Docker Registry Garbage Collector

![Docker Build Status](https://img.shields.io/docker/build/uphy/registry-garbage-collector.svg)
![Docker Automated build](https://img.shields.io/docker/automated/uphy/registry-garbage-collector.svg)

[Docker Registry](https://github.com/docker/distribution) has `registry garbage collect` command which removes all garbages.
The latest unreleased Docker Registry has one more option '-m' which removes also untagged images.
To use the latest feature, I've published this image.

This image uses latest `registry` command.
It removes the mounted /data directory.

## How to use

Collect garbages:

```bash
$ docker run -d -p "5000:5000" -v "$(pwd)/data/registry:/var/lib/registry" registry:2.6
$ docker pull bash
$ docker tag bash localhost:5000/bash
$ docker push localhost:5000/bash
$ docker run -i --rm -v "$(pwd)/data/registry:/target" uphy/registry-garbage-collector
```

Garbage collector server(with basic authentication):

```bash
docker run -i --rm \
  -p "8080:8080" \
  -v "$(pwd)/data/registry:/target" \
  uphy/registry-garbage-collector server
```

Garbage collector server(with basic authentication):

```bash
docker run -i --rm \
  -e "AUTH_USER=admin" \
  -e "AUTH_PASSWORD=password"  \
  -p "8080:8080" \
  -v "$(pwd)/data/registry:/target" \
  uphy/registry-garbage-collector server
```
# Docker Registry Garbage Collector

[Docker Registry](https://github.com/docker/distribution) has `registry garbage collect` command which removes all garbages.
The latest unreleased Docker Registry has one more option '-m' which removes also untagged images.
To use the latest feature, I've published this image.

This image uses latest `registry` command.
It removes the mounted /data directory.

## How to use

```bash
$ docker run -d -v "$(pwd)/data/registry:/var/lib/registry" registry:2.6
$ docker run -i --rm -v "$(pwd)/data/registry:/target" uphy/registry-garbage-collector
```

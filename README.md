# tuni-devops

Exercises for Tampere University course COMP.SE.140 Continuous Development and Deployment - DevOps

# Exercise 1 - Docker Compose

The purpose of this exercise is to learn (or recap) how to create a system of two interworking
services that are started up and stopped together.

## Usage

To build and start the containers and network, test they are working correctly, and finally remove the containers and network, run the following commands:

```bash
$ git clone -b exercise1 https://github.com/emilcalonius/tuni-devops.git
$ cd tuni-devops
$ docker compose up –-build -d # Wait for all containers to start
$ curl localhost:8199
$ docker compose down
```

# Author

Emil Calonius \
ec435282 \
emil.calonius@tuni.fi

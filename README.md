# Load Generator

This is a minimal web application with the following microservices

- **React frontend:** Uses react query to load data from an api and display the result.
- **Go APIs:** Both have `/` and `/ping` endpoints. `/` queries the Database for the current time and the number of requests for each api recorded within the database, and `/ping` returns `pong`.
- **Postgres Database:** An empty PostgreSQL database with no tables or data. Used to show how to set up connectivity. The API applications execute `SELECT NOW() as now;` to determine the current time to return.

We have also included a load tester to simulate client requests
- **Python Load Generator:** Queries the GO API at a configurable speed.

![](./docs/screenshot.png)


## Running the Application

While the whole point of this course is that you probably won't want/need to run the application locally, we can do so as a starting point.
- Please install Devbox according to their instructions: https://www.jetify.com/devbox/docs/installing_devbox/

```bash
# start the shell session
$ devbox shell
Starting a devbox shell...

# list the tasks
$ task --list-all

task: Available tasks for this project:
* bootstrap-buildx-builder:       Bootstrap the builder
* build:                          Build/push docker images with buildx (mutli-architecture)
* clean:                          Clean up all project resources
* clean_buildx:                   Remove buildx docker container, image, and volume
* clean_deep:                     Clean up all project resources (may affect others on your system)
* clean_images:                   Remove all images that are not being used by a container
* clean_volumes:                  Remove volumes from the project
* load-test:                      Run the load test
* prune_buildx:                   Remove all build cache from docker buildx
* run-local-registry:             Run a local registry
* run-psql-init-script:           Add tables to database
* start:                          Start all services with Docker Compose
* stop:                           Stop and remove all containers
* api-golang:build:               Build container image
* client:build:                   Build multi-arch container image
* load-generator:build:           Build multi-arch container image
* postgresql:build:               Build multi-arch container image
```

Here's a quick start

```bash
# get the apps going
$ t build
$ t start
$ t load-test


# and tear down
$ t stop
$ t clean_deep

```

---

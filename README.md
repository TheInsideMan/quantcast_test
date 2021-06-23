# Quantcast Test

https://github.com/TheInsideMan/quantcast_test

## Installation
1. Open a terminal and cd into the directory where this repo is located.
2. This project uses BuildKit so please run `$ export DOCKER_BUILDKIT=1`.
3. To build the docker container run `$ make`.
4. To use the new container `$ ./bin/cookie_parser -f cookie_log.csv -d 2018-12-09`.

## Unit Testing
If you want to unit test first `$ make unit-test` (note when working locally `make` must always be ran before `make unit-test`).

## CI Building
`.github/workflows/ci.yaml` contains a deployment process which runs `make unit-test` before proceeding. If any of the unit tests fail deployment is halted.
1. `make unit-test`
2. `make unit-test-coverage`
3. `make PLATFORM=some_platform`
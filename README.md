# simpals-backend-test-task

Test task for Simpals inverview on Backend engineer position

## Project structure

- ``worker`` - Web-worker. Collects data from ``data.json`` and sends them to the gRPC service (implemented)
- ``gRPC service`` - gRPC service. Collects incoming data and stores them into Elasticsearch (worker -> gRPC service = implemented; api -> gRPC service = implemented, not tested)
- ``API`` - API service. Requests data from Elasticsearch and returns them through GraphQL (Not implemented !!!!)

## Test coverage
- ``worker`` - Covered with tests, except gRPC related functionality
- ``gRPC service`` - Not covered with tests
- ``API`` - Not covered with tests

## Code documentation
Code not documented

## Launch

Simply run ``docker-compose up -d``
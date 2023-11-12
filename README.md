# simpals-backend-test-task

Test task for Simpals inverview on Backend engineer position

## Project structure

- ``worker`` - Web-worker. Collects data from ``data.json`` and sends them to the gRPC service
- ``service`` - gRPC service. Collects incoming data and stores them into Elasticsearch
- ``API`` - API service. Requests data from Elasticsearch and returns them through GraphQL
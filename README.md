# Arda

## Intentions
- Microservice
- Go
- PostgreSQL
- gRPC
- GraphQL
- API gateway
- TLS certificates
- OAuth, M2M, JWT
- Kafka
- Elasticsearch
- RabbitMQ
- Redis, Memcache?
- PCI DSS
- Cassandra
- MongoDB
- Docker
- Kubernetes
- Istio
- Security, Hardening
- Code Quality, Coverage
- Integration, E2E tests
- Stress, Soak test
- Monitoring, Tracing, OpenTelemetry, Performance, Alerting
- BigData, Data warehousing
- Chaos Monkey
- Fan-Out
- Websockets
- Audit Log
- PDFs, CSVs

## Services
| Service Name             | Service Purpose        |
| ------------------------ | ---------------------- |
| [mordor](mordor)         | checkout               |
| [rivendell](rivendell)   | inventory              |
| [gundabad](gundabad)     | supply                 |
| [weathertop](weathertop) | currency               |
| [gondor](gondor)         | cart                   |
| [angmar](angmar)         | recommendation         |
| [erebor](erebor)         | payments               |
| [esgaroth](esgaroth)     | gateway                |
| [arnor](arnor)           | push                   |
| [lorien](lorien)         | email                  |
| [osgiliath](osgiliath)   | notification           |
| [gorgoroth](gorgoroth)   | shipping               |
| [isengard](isengard)     | international delivery |
| [rohan](rohan)           | local delivery         |
| [moria](moria)           | ad                     |
| [orodruin](orodruin)     | profile                |
| [fangorn](fangorn)       | loadgen                |

## Installation instructions
- Install golang
```shell
brew install go
```
- Install protoc
```shell
brew install protobuf
```
- Install protoc-gen-go
```shell
go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

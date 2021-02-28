GO_DIRS := angmar arnor erebor esgaroth fangorn gondor gorgoroth gundabad isengard lorien mordor moria orodruin osgiliath rivendell rohan weathertop

COMPOSE := docker-compose -f build/local/docker-compose.yaml

update-vendor:
	for p in $(GO_DIRS); do (echo $$p; cd service/$$p && go mod tidy && go mod vendor); done

run:
	env \
		COMMAND="go run -mod=vendor cmd/serverd/main.go" \
		${COMPOSE} up

teardown:
	${COMPOSE} down -v

db:
	${COMPOSE} up -d db

DB_SERVICES := rivendell

migrate-up: db
migrate-up:
	sleep 2
	for p in $(DB_SERVICES); do (\
		echo $$p; \
		${COMPOSE} run --rm -v $(shell pwd)/service/$$p/data/migrations:/migrations db-migrate \
		sh -c './migrate -path /migrations -database postgres://arda:pwd@db:5432/$$p?sslmode=disable up'\
	); done

PROTO_SERVICES := rivendell

gen-proto:
	for p in $(PROTO_SERVICES); do (\
		echo $$p; \
		cd service/$$p && make gen-proto \
	); done

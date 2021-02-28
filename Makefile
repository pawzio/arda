GO_DIRS := angmar arnor erebor esgaroth fangorn gondor gorgoroth gundabad isengard lorien mordor moria orodruin osgiliath rivendell rohan weathertop

COMPOSE := docker-compose -f build/local/docker-compose.yaml

update-vendor:
	for p in $(GO_DIRS); do (echo $$p; cd $$p && go mod tidy && go mod vendor); done

run:
	env COMMAND="go run -mod=vendor cmd/serverd/main.go" ${COMPOSE} up

teardown:
	${COMPOSE} down -v

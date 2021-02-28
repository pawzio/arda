GO_DIRS := angmar arnor erebor esgaroth fangorn gondor gorgoroth gundabad isengard lorien mordor moria orodruin osgiliath rivendell rohan weathertop

update-vendor:
	for p in $(GO_DIRS); do (echo $$p; cd $$p && go mod tidy && go mod vendor); done

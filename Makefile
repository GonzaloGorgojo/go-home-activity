main_package_path = ./cmd/server/main.go
binary_name = homeserver

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


## run: runs migration and start app without build.
.PHONY: run
run:
	@echo "..."
	goose up && go run ./cmd/server/

## create-migration MIGRATION_NAME=<name>: creates a new migration file.
.PHONY: create-migration
create-migration:
	goose -s create $(MIGRATION_NAME) sql

## remove-migration: execute the down method on the last executed migration.
.PHONY: remove-migration
remove-migration:
	goose down

## audit: run quality control checks
.PHONY: audit
audit:
		go mod tidy -diff
		go mod verify
		test -z "$(shell gofmt -l .)" 
		go vet ./...
		go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
		go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## build: build the application and place it in tmp
.PHONY: build
build:
	go build -o=/tmp/bin/${binary_name} ${main_package_path}

## run-build: build and runs the app.
.PHONY: run-build
run-build: build
	/tmp/bin/${binary_name}
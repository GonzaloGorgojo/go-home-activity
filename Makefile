run:
	@echo "Running migration and starting app..."
	goose up && go run ./cmd/server/

check:
	@echo "Making requests..."
	 ./requests.sh

create-migration:
	@echo "Creating migration file with name: $(MIGRATION_NAME)"
	goose -s create $(MIGRATION_NAME) sql

remove-migration:
	@echo "Removing last migration"
	goose down

audit:
		@echo "Running audit..."
		go mod tidy -diff
		go mod verify
		test -z "$(shell gofmt -l .)" 
		go vet ./...
		go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
		go run golang.org/x/vuln/cmd/govulncheck@latest ./...
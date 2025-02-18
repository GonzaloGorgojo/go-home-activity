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
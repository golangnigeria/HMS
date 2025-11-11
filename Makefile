# Windows-safe Goose Makefile (PowerShell version)
GOOSE = $(GOPATH)/bin/goose
MIGRATIONS_DIR = internals/driver/migrations

# Create new migration
new:
	if not exist $(MIGRATIONS_DIR) mkdir $(MIGRATIONS_DIR)
	powershell -Command "$$env:DATABASE_URL=(Get-Content .env | ForEach-Object { if ($$_ -match '^DATABASE_URL=(.*)$$') { $$matches[1] } }); & '$(GOOSE)' -dir $(MIGRATIONS_DIR) create $(name) sql"

# Apply all up migrations
migrate-up:
	powershell -Command "$$env:DATABASE_URL=(Get-Content .env | ForEach-Object { if ($$_ -match '^DATABASE_URL=(.*)$$') { $$matches[1] } }); & '$(GOOSE)' -dir $(MIGRATIONS_DIR) postgres $$env:DATABASE_URL up"

# Roll back one migration
migrate-down:
	powershell -Command "$$env:DATABASE_URL=(Get-Content .env | ForEach-Object { if ($$_ -match '^DATABASE_URL=(.*)$$') { $$matches[1] } }); & '$(GOOSE)' -dir $(MIGRATIONS_DIR) postgres $$env:DATABASE_URL down"

# Show migration status
migrate-status:
	powershell -Command "$$env:DATABASE_URL=(Get-Content .env | ForEach-Object { if ($$_ -match '^DATABASE_URL=(.*)$$') { $$matches[1] } }); & '$(GOOSE)' -dir $(MIGRATIONS_DIR) postgres $$env:DATABASE_URL status"

# run go file
run:
	go run ./cmd/web

# run tailwind file
compile:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch  
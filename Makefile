run:
	air -c .air.toml


.SILENT:
.PHONY: sql
sql: 
	psql $(shell cat .env | grep DATABASE_URL | cut -d '=' -f2)

stage: 
	templ generate&& tailwindcss -i ./static/input.css -o ./static/output.css

migrate-up:
	migrate -database $(shell cat .env | grep DATABASE_URL | cut -d '=' -f2) -path db/migration up

migrate-down:
	migrate -database $(shell cat .env | grep DATABASE_URL | cut -d '=' -f2) -path db/migration down


mock-migrate-up:
	migrate -database $(shell cat .env | grep MOCK_DATABASE | cut -d '=' -f2) -path db/migration up

mock-migrate-down:
	migrate -database $(shell cat .env | grep MOCK_DATABASE | cut -d '=' -f2) -path db/migration down

run:
	air -c .air.toml

css:
	tailwindcss -i ./static/input.css -o ./static/output.css --watch

.SILENT:
.PHONY: sql
sql: 
	psql $(shell cat .env | grep DATABASE_NEW | cut -d '=' -f2)

stage: 
	templ generate&& tailwindcss -i ./static/input.css -o ./static/output.css

migrate-up:
	migrate -database $(shell cat .env | grep DATABASE_NEW | cut -d '=' -f2) -path migration up

migrate-down:
	migrate -database $(shell cat .env | grep DATABASE_NEW | cut -d '=' -f2) -path migration down

run:
	air -c .air.toml

css:
	tailwindcss -i ./static/input.css -o ./static/output.css --watch

.SILENT:
.PHONY: sql
sql: 
	psql $(shell cat .env | grep DATABASE_URL | cut -d '=' -f2)

stage: 
	tailwindcss -i ./static/input.css -o ./static/output.css && templ generate

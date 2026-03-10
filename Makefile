include .env
DB_DSN=postgres://$(MAIN_DB_USER):$(MAIN_DB_PASSWORD)@$(MAIN_DB_HOST):$(MAIN_DB_PORT)/$(MAIN_DB)?sslmode=$(MAIN_DB_SSLMODE)

create-schema:
	go run -mod=mod entgo.io/ent/cmd/ent new $(name)

generate-schema:
	go generate ./ent

create-migration:
	atlas migrate diff $(name) --dir "file://ent/migrate/migrations" --to "ent://ent/schema" --dev-url "docker://postgres/17.2/ent"

apply-migration:
	atlas migrate apply --dir "file://ent/migrate/migrations" --url $(DB_DSN)

run-dev:
	air
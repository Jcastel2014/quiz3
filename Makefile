include .envrc

.PHONY: run
run:
	@echo  'Running application…'
	@go run ./cmd/api

.PHONY: post
post:
	@BODY='{"email":"javiercaste94@gmail.com", "firstName":"Javier", "middleName":"Francisco", "lastName":"Castellanos"}'; \
	echo "$$BODY"; \
	curl -i -d "$$BODY" localhost:4000/signUp

.PHONY: update
update:
	@UPDATE='{"email":"${email}", "row":${row}}'; \
	echo "$$UPDATE"; \
	curl -X PATCH localhost:4000/signUp/update -d "$$UPDATE"

.PHONY: delete
delete:
	@DELETE='{"delete":${del}}'; \
	echo "$$DELETE"; \
	curl -X DELETE localhost:4000/signUp/delete -d "$$DELETE"
.PHONY: read
read:
	@curl -i localhost:4000/signUp/read

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@echo 'Running application...'
	@go run ./cmd/api -port=4000 -env=development -db-dsn=${COMMENTS_DB_DSN}

## db/psql : connect to the database using psql (terminal)

.PHONY: db/psql
db/psql:
	psql ${QUIZ3_DB_DSN}


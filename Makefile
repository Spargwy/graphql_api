ifneq (,$(wildcard ./.env))
    include .env
    export
endif

docker-run:
	docker-compose up -d --build

docker-migrate:
	docker run --rm --network=graphql_api_dbmate -v "$(shell pwd)/db:/db"\
	 -e DATABASE_URL="${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}"\
		amacneil/dbmate:latest up

docker-test-data:
	docker-compose exec db psql -f app/testData/seed.sql ${DB_NAME} -U ${DB_USER}

migrate:
	dbmate --url ${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE} up

local-test-data:
	cat testDB/seed.sql | psql ${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}

lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.43.0 golangci-lint run -v
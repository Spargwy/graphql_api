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
	docker-compose exec db cat app/seed.sql | psql ${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}

local-test-data:
	cat seed.sql | psql ${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}
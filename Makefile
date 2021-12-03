ifneq (,$(wildcard ./.env))
    include .env
    export
endif

docker-run:
	docker-compose up -d --build

docker-migrate:
	docker run --rm --network=graphql_api_dbmate -v "$(shell pwd)/db:/db"\
	 -e DATABASE_URL="postgres://${DB_USER}:${DB_PASSWORD}@postgres/${DB_NAME}?sslmode=${DB_SSLMODE}"\
		amacneil/dbmate:latest up

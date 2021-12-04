ifneq (,$(wildcard ./.env))
    include .env
    export
endif

docker-run:
	docker-compose up -d --build

docker-migrate:
	docker run --rm --network=graphql_api_dbmate -v "$(shell pwd)/db:/db"\
	 -e DATABASE_URL="postgres://postgres:postgres@postgres:5432/products?sslmode=disable"\
		amacneil/dbmate:latest 

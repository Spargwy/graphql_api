docker-run:
	docker-compose up -d --build

docker-logs:
	docker-compose logs -f web
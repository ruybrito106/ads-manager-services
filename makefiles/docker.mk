docker-compose/restart:
	docker-compose down
	docker-compose up -d --build

restart: docker-compose/restart bootstrap
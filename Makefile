up:
	docker-compose -f dockers/docker-compose.yml build && docker-compose -f dockers/docker-compose.yml up -d
down:
	docker-compose -f dockers/docker-compose.yml down
logs:
	docker-compose -f dockers/docker-compose.yml logs -f app
test:
	docker-compose -f dockers/docker-compose.yml exec app go test ./tests/...
wire:
	cd package/infrastructure/di && wire

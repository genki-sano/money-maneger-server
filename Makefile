up:
	docker-compose -f dockers/docker-compose.yml build && docker-compose -f dockers/docker-compose.yml up
down:
	docker-compose -f dockers/docker-compose.yml down
wire:
	cd package/infrastructure/di && wire

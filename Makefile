up:
	docker-compose -f dockers/docker-compose.yml build && docker-compose -f dockers/docker-compose.yml up

lint:
	golint ./...

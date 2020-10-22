run:
	go run main.go web

up:
	docker-compose -f dockers/docker-compose.yml build && docker-compose -f dockers/docker-compose.yml up

lint:
	golint ./...

deploy:
	git push heroku master

config:
	heroku config

open:
	heroku open

logs:
	heroku logs --tail

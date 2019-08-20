
usage:
	@echo "build up stop run test"

build:
	docker-compose build

up:
	docker-compose up

stop:
	docker-compose stop

run:
	docker-compose run --rm app bundle exec go run

test:
	docker-compose run --rm app bundle exec go test

build:
	docker-compose build
up:
	docker-compose up
run:
  docker-compose run --rm app bundle exec go run
test:
  docker-compose run --rm app bundle exec go test

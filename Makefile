build:
	docker-compose build challenge

run:
	docker-compose up challenge

migrate:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres' up
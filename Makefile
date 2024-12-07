serve:
	cd api && go run main.go serve

types:
	cd api && go run main.go ts

entgo:
	cd api && go generate ./...

docker:
	docker compose -f ./docker-compose-dev.yml up --build --force-recreate
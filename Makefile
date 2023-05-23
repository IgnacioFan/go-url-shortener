-include .env
export

app.build:
	docker-compose up --build

app.start:
	docker-compose up

app.stop:
	docker compose down

db.cli:
	docker exec -it $(POSTGRES_HOST) psql -U $(POSTGRES_USER)

gen.env:
	./scripts/generate_env.sh

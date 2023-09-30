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

test.prep:
	docker-compose -f docker-compose.test.yml up --build -d

gen.env:
	./scripts/generate_env.sh

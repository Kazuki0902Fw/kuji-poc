.PHONY: up down build logs migrate-new migrate-up migrate-down migrate-down-all migrate-status bobgen-mysql gqlgen go-get register-test-data

# Docker Compose操作
up:
	docker-compose up -d

down:
	docker-compose down -v

build:
	docker-compose build --no-cache

logs:
	docker-compose logs -f

# マイグレーション
migrate-new:
	docker-compose exec app sql-migrate new ${NAME}

migrate-up:
	docker-compose exec app sql-migrate up

migrate-down:
	docker-compose exec app sql-migrate down

migrate-down-all:
	docker-compose exec app sql-migrate down --limit=0

migrate-status:
	docker-compose exec app sql-migrate status

bobgen-mysql:
	docker-compose exec app ash -c "cd /src && bobgen-mysql -c bobgen.yml"

gqlgen:
	docker-compose exec app ash -c "cd /src && gqlgen"

go-get:
	docker-compose exec app ash -c "cd /src && go get ${PKG}"

register-test-data:
	docker-compose exec app ash -c "./scripts/register_test_data.sh"

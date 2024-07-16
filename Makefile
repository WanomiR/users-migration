

read-env:
	@. ./read-env.sh

migration_create:
	@migrate create -ext sql -dir postgres/migration/ -seq init_mg

migration_up:
	@migrate -path postgres/migration/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migration_down:
	@migrate -path postgres/migration/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

migration_fix:
	@migrate -path postgres/migration/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" force $(VERSION)

postgres_dump:
	@docker exec -t $(CONT_NAME) pg_dump -U $(DB_USER) -h $(DB_HOST) -d $(DB_NAME) > ./postgres/create_tables.sql
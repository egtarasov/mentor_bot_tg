POSTGRES_SETUP_TEST := user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} host=localhost port=5432 sslmode=disable


migration-up:
	goose -dir ./migrations postgres "${POSTGRES_SETUP_TEST}" up

migration-down:
	goose -dir ./migrations postgres "${POSTGRES_SETUP_TEST}" reset


build:
	go build -o app ./cmd/main.go
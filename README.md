# How to deploy the bot?

1. Go to `deploy/.env`. Enter the correct values for telegram-token and postgresql database settings (TELEGRAM_TOKEN, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_HOST, POSTGRES_PORT, POSTGRES_SSLMODE)
2. Write the sql script to insert info about commands, materials, pictures, goals. Alternatively, you can use [goose](https://github.com/pressly/goose) to write  migration and put it into `app/migrations`.
3. Enter `make deploy` or `docker-compose up --build`. If needed, you can modify the `docker-compose.yaml`.
4. Test the bot and inspect containers logs if something went wrong.
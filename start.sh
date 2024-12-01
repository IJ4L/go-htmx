#!/bin/sh

set -e

echo "run db migration"

ls -la /app

if [ -f /app/app.env ]; then
    source /app/.env
fi

/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
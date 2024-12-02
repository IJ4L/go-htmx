#!/bin/sh

set -e

echo "run db migration"

if [ -f /app/app.env ]; then
    source /app/.env
fi

/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
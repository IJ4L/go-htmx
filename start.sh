#!/bin/sh

set -e

echo "run db migration"

ls -l

if [ -f /app/app.env ]; then
    source /app/app.env
    echo "source /app/app.env"
else
    echo "/app/app.env does not exist"
fi

echo "DB_SOURCE: $DB_SOURCE"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
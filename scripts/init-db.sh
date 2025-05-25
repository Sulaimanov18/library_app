# #!/bin/bash
# set -e

# # Wait for PostgreSQL to be ready
# until PGPASSWORD=$POSTGRES_PASSWORD psql -h "$DB_HOST" -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c '\q'; do
#   echo "PostgreSQL is unavailable - sleeping"
#   sleep 1
# done

# echo "PostgreSQL is up - executing migrations"

# # Run migrations
# cd /app
# goose -dir migrations postgres "host=$DB_HOST user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB sslmode=disable" up

# echo "Migrations completed" 
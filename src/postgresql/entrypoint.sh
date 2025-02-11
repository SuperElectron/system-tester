#!/bin/sh

set -e  # Exit on error

echo "Waiting for PostgreSQL to start..."
until pg_isready -h localhost -U "$POSTGRES_USER"; do
  sleep 1
done

echo "PostgreSQL is ready. Running migrations..."
psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f /docker-entrypoint-initdb.d/000001_create_users_table.up.sql

echo "Starting PostgreSQL..."
exec postgres

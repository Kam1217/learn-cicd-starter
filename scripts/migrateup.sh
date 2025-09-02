#!/bin/bash

if [ -f .env ]; then
    source .env
fi

cd sql/schema
goose -dir sql/schema sqlite3 $DATABASE_URL up

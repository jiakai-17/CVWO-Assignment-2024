FROM postgres:15-alpine

# Set up the database and seed it with starting data
COPY schema.sql /docker-entrypoint-initdb.d/01-schema.sql
COPY triggers.sql /docker-entrypoint-initdb.d/02-triggers.sql
COPY seed.sql /docker-entrypoint-initdb.d/03-seed.sql

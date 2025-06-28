FROM postgres:17-alpine
LABEL maintainer="Adrien Panis <a.panis@my-goodlife.fr>"
LABEL description="PostgreSQL 17 with custom configuration for My Good Life"
LABEL version="1.0"

# Copy initialization script
COPY build-docker/database/init.sql /docker-entrypoint-initdb.d/

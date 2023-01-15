FROM jbergknoff/postgresql-client

WORKDIR /usr/src/app

COPY ./init .

ENTRYPOINT psql postgresql://$(echo $POSTGRES_USER):$(echo $POSTGRES_PASSWORD)@$(echo $POSTGRES_HOST):$(echo $POSTGRES_PORT)/$(echo $POSTGRES_DB) -a -f init.sql

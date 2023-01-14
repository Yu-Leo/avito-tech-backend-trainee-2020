FROM jbergknoff/postgresql-client

WORKDIR /usr/src/app

COPY ./init .

ENTRYPOINT ["psql", "postgresql://postgres:postgres@postgres:5432/postgres", "-a", "-f", "init.sql"]

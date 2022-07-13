# Evresi: API

## Running
The environment variable `POSTGRES_URL` must be set with the URL to be used for connecting to the PostgreSQL database.
See the below section for a convenient way to set this variable during development.

## Developing
The `make run` target will grab environment variables from a file named `.env` if it exists.

The following can be used as a template for setting the `POSTGRES_URL` variable within `.env`:
```
POSTGRES_URL=postgres://username:password@localhost:5432/evresi
```

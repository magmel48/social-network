# social-network

In case of any changes in API spec run:
```shell
oapi-codegen -package api api/openapi.json > internal/api/api.gen.go
```

Before first run:
```shell
cp .env.example .env
```
and fill with the values.

For local development run:
```shell
docker-compose up api
```

For local migrations change `MYSQL_PORT` and run:
```shell
make migrate up
```

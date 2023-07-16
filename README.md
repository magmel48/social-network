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

Known issues:
- https://github.com/kyleconroy/sqlc/issues/695
- `ON DUPLICATE KEY UPDATE` is better than `INGORE`, an `IGNORE` will really ignore any error, not only unique constraint violation.

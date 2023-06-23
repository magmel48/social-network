# social-network

In case of any changes in API spec:
```shell
oapi-codegen -package api api/openapi.json > internal/api/api.gen.go
```

Before first run:
```shell
cp .env.example .env
```

For local debug:
```shell
docker-compose up api
```

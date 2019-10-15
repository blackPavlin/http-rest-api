# http-rest-api


### Database URL

Production
```
postgres://postgres:123@localhost:5432/restapi_dev?sslmode=disable
```

Test
```
postgres://postgres:123@localhost:5432/restapi_test?sslmode=disable
```

### Migrations
```
migrate -path migrations -database "postgres://postgres:123@localhost:5432/restapi_dev?sslmode=disable" up
```
```
migrate -path migrations -database "postgres://postgres:123@localhost:5432/restapi_dev?sslmode=disable" down
```

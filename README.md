# gin-gonic.com-docs-examples-only-bind-query-string

- Only bind query string

- Reference: https://gin-gonic.com/docs/examples/only-bind-query-string/

## gvm

```sh
gvm install go1.23.5
gvm use go1.23.5
```

## go get

```sh
go get .
```

## go run

```sh
go run .
```

## cURL

- GET /person

```sh
curl --location --request GET 'localhost:8080/person?name=NAME&address=ADDRESS' \
--form 'message="Hello"' \
--form 'name="admin"'
```

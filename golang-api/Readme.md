
# Golang API

REST API example.

## Deploy

```
$ up
```

## Usage

```
$ curl -d Tobi https://example.com/
welcome to the family "Tobi"!

$ curl -d Loki https://example.com/
welcome to the family "Loki"!

$ curl -d Jane https://example.com/
welcome to the family "Jane"!

$ curl https://example.com/
- Tobi
- Jane
- Loki
```

## Notes

There's no need to wrap `pets` in a mutex since
Lambda processes only a single request per container
concurrently.

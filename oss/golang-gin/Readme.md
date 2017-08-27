
# Golang Gin Service


## Install dependencies

This will install the required dependencies and uses the [Go Dependency Management Tool](https://github.com/golang/dep),
which you may need to install.

```
$ make deps
```

## Run Locally

You can also run this up locally using the following

```
$ make build
$ PORT=3000 ./server
```

## Deploy

This will deploy to default `AWS` profile.

```
$ make deploy
```
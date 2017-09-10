
# Node Secrets

Environment variable management via encrypted secrets.

## Setup

First define a few fake environment variables, they are encrypted by default,
and contain anything you like, even JSON for example.

```
$ up env add MSG Hello
$ up env add NAME Tobi
```

By default env variables are scoped to _all_ stages, however any defined
for a specific stage take precedence. Let's define one for production with the `-s` or `--stage` flag.

```
$ up env add MSG "Hello from production" -s production
```

## Deploy

Deploy development:

```
$ up
```

Deploy production:

```
$ up deploy production
```

View both:

```
$ up url -o
$ up url production -o
```


# Node Secrets

Environment variable management via encrypted secrets.

## Deploy

First define a few fake environment variables, they are encrypted by default,
and contain anything you like, even JSON for example.

```
$ up env add MSG Hello
$ up env add NAME Tobi
```

List the env vars with:

```
$ up env
```

Deploy the environments:

```
$ up
$ up deploy staging
$ up deploy production
```

View them:

```
$ up url -o
$ up url -o staging
$ up url -o production
```

By default env variables are scoped to _all_ stages, so you'll see the same results. You may also define env vars scoped to a particular stage, using the `-s` or `--stage` flag:

```
$ up env add -o staging MSG NAME Loki
$ up env add -o production MSG NAME Jane
```

Now restart by deploying again:

```
$ up deploy staging
$ up deploy production
```

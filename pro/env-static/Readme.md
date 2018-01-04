
# Environment Static

Environment variable availability during hooks.

## Deploy

Map the `NAME` environment variable to different names, based on the stage.

```
$ up env set NAME Tobi
$ up env set -s staging NAME Loki
u$ p env set -s production NAME Jane
```

Deploy each stage:

```
$ up production
$ up staging
$ up
```

View each response:

```
$ curl -s `up url`
# contains: Hello Tobi from development

$ curl -s `up url staging`
# contains: Hello Loki from staging

$ curl -s `up url production`
# contains: Hello Jane from production
```

## Notes

This example illustrates how environment variables are available at build time,
so you may use `UP_STAGE`, `NODE_ENV` which is implied, or any custom env vars
mapped to stages via `up env` to alter your builds on these values.

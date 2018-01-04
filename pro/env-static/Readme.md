
# Environment Static

Environment variable availability during hooks.

## Deploy

Map the `NAME` environment variable to different names, based on the stage.

```
$ up env set NAME Tobi
$ up env set -s staging NAME Loki
$ up env set -s production NAME Jane
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

This example illustrates how environment variables are available at build time, so you may utilize vars mapped to the stage you are deploying to alter configuration, client-side builds, and so on. Note that `UP_STAGE` and `NODE_ENV` by default will be the target stage name (`development`, `staging`, `production`).

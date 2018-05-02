
# Rollbacks

Rollback to a previous production release.

## Deploy

```
$ up
```

## Notes

Deploy the app to production a few times:

```
$ up production
```

Check out the version:

```
$ curl `up url -s production`
Version 5
```

Rollback:

```
$ up rollback
$ curl `up url -s production`
Version 4
```

Rollback to a specific version:

```
$ up rollback 5
$ curl `up url -s production`
Version 5
```

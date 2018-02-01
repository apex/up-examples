
# Node Debug

Logging using stderr example.

## Deploy

```
$ up
```

## Notes

By default Up treats STDERR output as error logs, this is not ideal in some cases such as Node's `debug()` module, which may then trigger alerts if you enable it in production.

Tuning the log settings in `up.json` can disable this default behaviour.

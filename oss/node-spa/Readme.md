
# Node SPA

Single page app using React & Neutrino.

## Deploy

```
$ up
```

## Notes

The `.static.dir` setting points Up to the Neutrino build output, which is created by `.hooks.build`,
and remove after deploy via `.hooks.clean`.

A rewrite of `/*` to `/` is used so that files which do not exist are routed to the index file,
allowing users to request for example the "/admin/settings" page and load "/index.html" with
your JavaScript.

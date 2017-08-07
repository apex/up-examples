
# Rewrite

Rewrites and redirection.

## Deploy

```
$ up
```

## Notes

Rewriting and redirection work with non-static apps and APIs as well.

The `/*` rule is a rewrite, meaning it does not perform an HTTP redirect, it replaces the path while serving, so that any path (`/admin/settings`, `/admin`, etc) will respond with `/` (index.html). This is ideal for single-page apps.

The `/store/:brand/products/:product"` rule matches the path segments on the left, and replaces with the `location` defined, performing a 302 Found redirect.

The `/blog` rule matches only that exact path, and redirects to an entirely separate host.

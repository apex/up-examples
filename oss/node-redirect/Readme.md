
# Node Redirect

Tiny redirect app.

## Deploy

```
$ up
```

## Notes

Typically you'd want to redirect a `www` domain – such as `www.example.com` – to `example.com`,
or vice-versa in order to reduce duplicate content. For now Up does not magically do this for you,
however you can deploy this tiny app.

With Up Pro simply run:

```
$ up env set REDIRECT_URL https://example.com
```

Change the `domain` in up.json and run:

```
$ up
```

Run the following to set the redirect:

```
$ curl -v `up url`
```

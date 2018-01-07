
# Golang CORS SPA

Golang single page app and CORS API.

## Setup

This example uses [Parcel](https://parceljs.org) for the SPA.

```
$ yarn global add parcel-bundler
```

## Deploy

Deploy the api and app:

```
$ up -C api
$ up -C app
```

We need `api.js`'s `process.env.API_URL` to point to a different url in production or staging, so
we define `API_URL` with Up Pro's `up env` command, pointing it to the API's endpoint.

```
$ up -C app env set API_URL $(up -C api url)
```

This may look a bit confusing, but it's effectively the same as the following, `-C` simply does the `cd` for you.

```
$ api=$(up -C api url)
$ cd app
$ up env set API_URL $api
```

Note that `process.env.API_URL` is replaced at build time, as these environment variables are available to hooks. Finally
open the app's url in the browser to check it out:

```
$ up -C app url -o
```

## Notes

- We use Parcel's `--public-url .` to remove the need for a `/build/` asset prefix
- We point up.json `static.dir` to the `./build` directory
- We `.upignore` `./src` and `./node_modules` as we only need `./build`

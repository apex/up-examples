# Node Micro

Run the [`micro`](https://github.com/zeit/micro) binary.

This makes us of `async`/`await` and demonstrates a GraphQL query (bonus).

For an example using `micro` programmatically and/or without a custom Node binary, check out the [`node-micro-alt` example](https://github.com/apex/up-examples/blob/master/oss/node-micro-alt).

## Setup

```
$ yarn
```

## Deploy

```
$ up
```

## Notes

The `proxy.command` script is run inside Lambda to start your server, used here to specify the host & port values for `micro`.

> **Note**: Micro's default `host` value is `::`, which is not compatible with Lambda; using `localhost` works, though! Also note that the micro(1) binary itself does literally nothing other than start your server, so node-micro-alt would be more idioimatic Node.js code.

```json
{
  "proxy": {
    "command": "./node_modules/.bin/micro --host localhost --port $PORT"
  }
}
```

# Node Micro

Run the [`micro`](https://github.com/zeit/micro) binary with a non-native Node v8.x runtime.

> This example builds upon the [`node-8` example](https://github.com/apex/up-examples/blob/master/oss/node-8).

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

The `build` hook in `up.json` simply runs `make`, which downloads the `node-v8.9.0-linux-x64` into the root directory.

```json
{
  "hooks": {
    "build": "make"
  }
}
```

The `proxy.command` script is run inside Lambda to start your server.

We use our "custom" `node` binary (`node-v8.9.0-linux-x64`) to run the `micro` binary, specifying the host & port values.

> **Note**: Micro's default `host` value is `::`, which is not compatible with Lambda; using `localhost` works, though! Also note that the micro(1) binary itself does literally nothing other than start your server, so node-micro-alt would be more idioimatic Node.js code.

```json
{
  "proxy": {
    "command": "./node-v8.9.0-linux-x64/bin/node node_modules/.bin/micro --host localhost --port $PORT"
  }
}
```

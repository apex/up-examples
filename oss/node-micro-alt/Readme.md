# Node Micro

Run the [`micro`](https://github.com/zeit/micro) programmatically.

This makes us of `async`/`await` and demonstrates a GraphQL query (bonus).

For an example using the `micro` binary and/or with a custom Node binary, check out the [`node-micro` example](https://github.com/apex/up-examples/blob/master/oss/node-micro).

## Setup

```
$ yarn
```

## Deploy

```
$ up
```

## Notes

We use the [`fast-async`](https://github.com/MatAtBread/fast-async) Babel plugin (based on [`nodent`](https://github.com/MatAtBread/nodent)) to compile away the `async` & `await` syntax into Promises. This is because Lambda supports Node v6.x (LTS), which does not support the syntax directly.

Lastly, the compiled output (`lib` directory) is excluded from version control (via `.gitignore`) but is re-included to the Up instance (via `.upignore`).

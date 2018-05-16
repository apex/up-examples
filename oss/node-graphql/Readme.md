
# Node GraphQL

GraphQL server with CORS.

## Setup

```
$ yarn
```

## Deploy

```
$ up
```

## Usage

With curl:

```
$ curl -d '{ "query": "{ pet(id: 0) { name }}" }' `up url`?pretty
```

With `fetch()` in the browser:

``js
const body = JSON.stringify({
  query: `query {
    pet(id: 2) {
      name
    }
  }`
})

const res = await fetch('http://localhost:3000', {
  headers: { 'Content-Type': 'application/json', Accept: 'application/json' },
  method: 'POST',
  body
})

console.log(res)
console.log(await res.json())
```

See the [CORS](https://up.docs.apex.sh/#configuration.cross_origin_resource_sharing) second of the documentation for details.

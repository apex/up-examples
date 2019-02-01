
# Node Express

Express application using the ncc bundler.

## Setup

```
$ yarn
$ npm i -g @zeit/ncc
```

## Deploy

```
$ up
```

## Notes

Create an `.upignore` file to omit everything except our package.json and ./server files:

```
*
!package.json
!server/**
```

The Up proxy command in up.json is changed to `node server/index.js` instead of the default `node app.js`,
allowing staging and production to run the build.

## Links

- [ncc](https://github.com/zeit/ncc) bundler alternative to Webpack

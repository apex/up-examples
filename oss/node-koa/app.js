const serve = require('koa-static')
const Koa = require('koa')
const compress = require('koa-compress')
const app = new Koa

const { PORT = 3000 } = process.env

// Uncomment this middleware to enable origin compression. This means
// Koa will perform the compression instead of Up, either techinque
// works fine.

// app.use(compress())

app.use(serve('public'))
app.listen(PORT)

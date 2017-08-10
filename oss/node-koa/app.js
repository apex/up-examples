const serve = require('koa-static')
const Koa = require('koa')
const app = new Koa

const { PORT = 3000 } = process.env
app.use(serve('public'))
app.listen(PORT)

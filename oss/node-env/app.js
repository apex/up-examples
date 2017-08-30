const http = require('http')

const { PORT = 3000 } = process.env

const json = (res, v) => {
  res.setHeader('Content-Type', 'application/json')
  res.end(JSON.stringify(v, null, 2))
}

http.createServer((req, res) => {
  json(res, {
    FOO: process.env.FOO,
    BAR: process.env.BAR,
  })
}).listen(PORT)

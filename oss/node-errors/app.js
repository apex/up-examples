const http = require('http')
const { PORT = 3000 } = process.env

http.createServer((req, res) => {
  if (Math.random() > .90) {
    res.statusCode = 500 + (Math.random() * 10 | 0)
    res.end('5xx error')
    console.log('%s %s -> %d', req.method, req.url, res.statusCode)
    return
  }

  if (Math.random() > .75) {
    res.statusCode = 400 + (Math.random() * 10 | 0)
    res.end('4xx error')
    console.log('%s %s -> %d', req.method, req.url, res.statusCode)
    return
  }

  res.end('ok')
  console.log('%s %s -> %d', req.method, req.url, 200)
}).listen(PORT)

const http = require('http')
const { PORT = 3000 } = process.env

http.createServer((req, res) => {
  console.log('%s %s', req.method, req.url)
  for (let k in req.headers) {
    console.log('  - %s: %s', k, req.headers[k])
  }
  res.end('Hello\n')
}).listen(PORT)

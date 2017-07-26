const http = require('http')

const port = parseInt(process.env.PORT || "3000", 10)

http.createServer((req, res) => {
  console.log('%s %s', req.method, req.url)
  res.end('Hello World\n')
}).listen(port, '127.0.0.1')

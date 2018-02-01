const http = require('http')
const debug = require('debug')('myapp')
const { PORT = 3000 } = process.env

http.createServer((req, res) => {
  debug('%s %s', req.method, req.url)
  debug('doing some stuff')
  debug('doing more stuff here')
  res.end('Hello World from Node.js\n')
}).listen(PORT)

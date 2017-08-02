const http = require('http')
const port = ~~process.env.PORT

http.createServer((req, res) => {
  res.end('Hello World from Node.js\n')
}).listen(port)

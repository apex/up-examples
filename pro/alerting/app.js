const http = require('http')
const { PORT = 3000 } = process.env

http.createServer((req, res) => {
  if (Math.random() > .75) {
    res.statusCode = 500 + (Math.random() * 10 | 0)
    res.end('5xx error')
    return
  }

  if (Math.random() > .50) {
    res.statusCode = 400 + (Math.random() * 10 | 0)
    res.end('4xx error')
    return
  }

  res.end('ok')
}).listen(PORT)

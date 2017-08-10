const http = require('http')
const { PORT = 3000 } = process.env

http.createServer((req, res) => {
  if (Math.random() > .90) {
    asdf()
  } else {
    res.end('Hello World')
  }
}).listen(PORT)

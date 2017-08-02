const http = require('http')
const port = ~~process.env.PORT

http.createServer((req, res) => {
  if (Math.random() > .90) {
    asdf()
  } else {
    res.end('Hello World')
  }
}).listen(port)

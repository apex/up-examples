const http = require('http')
const { PORT = 3000, REDIRECT_URL } = process.env

http.createServer((req, res) => {
  res.setHeader('Content-Type', 'text/plain')
  res.setHeader('Location', REDIRECT_URL)
  res.statusCode = 301
  res.end('Redirecting to ' + REDIRECT_URL)
}).listen(PORT)

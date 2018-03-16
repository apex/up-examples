const http = require('http')
const { PORT = 3000 } = process.env

http.createServer((req, res) => {
  res.setHeader('Content-Type', 'text/html; charset=utf-8')
  res.end(`<!DOCTYPE html>
  <html>
    <head>
      <meta charset="utf-8">
      <title>Node Static</title>
      <link rel="stylesheet" href="/style.css">
    </head>
    <body>
      <p>Hello World</p>
    </body>
  </html>`)
}).listen(PORT)

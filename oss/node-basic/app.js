const http = require('http')
const { PORT = 3000, AWS_EXECUTION_ENV } = process.env

http.createServer((req, res) => {
  res.end(`Hello World from ${AWS_EXECUTION_ENV}`)
}).listen(PORT)

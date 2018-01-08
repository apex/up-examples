const http = require('http')
const { PORT = 3000 } = process.env

http.createServer((req, res) => {
  const mem = process.env.AWS_LAMBDA_FUNCTION_MEMORY_SIZE
  res.end('Memory: ' + mem + '\n')
}).listen(PORT)

const http = require('http')
const { PORT = 3000 } = process.env
const version = process.env.AWS_LAMBDA_FUNCTION_VERSION

http.createServer((req, res) => {
  res.end('Version ' + version + '\n')
}).listen(PORT)

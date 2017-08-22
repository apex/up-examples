const sleep = require('sleep-promise')
const http = require('http')

const { PORT = 3000 } = process.env

http.createServer(async (req, res) => {
  await sleep(1)
  res.end('Hello World from Node 8\n')
}).listen(PORT)

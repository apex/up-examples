const http = require('http')
const sleep = require('sleep-promise')

const { PORT = 3000 } = process.env
const sleepTime = 10

http.createServer(async (req, res) => {
  await sleep(sleepTime)
  res.end('Hello World from Node 8\n')
}).listen(PORT)
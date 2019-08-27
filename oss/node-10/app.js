const axios = require('axios')
const http = require('http')

const { PORT = 3000 } = process.env
const url = 'https://apex.sh'

http.createServer(async (req, res) => {
  const start = Date.now()
  await axios.get(url)
  res.end(`Response time: ${Date.now() - start}ms\nNode version: ${process.version}`)
}).listen(PORT)

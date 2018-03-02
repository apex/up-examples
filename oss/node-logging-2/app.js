const https = require('https')
const http = require('http')
const { Logger } = require('up')
const { PORT = 3000 } = process.env

const log = new Logger

http.createServer((req, res) => {
  const url = 'https://apex.sh'
  const start = Date.now()

  log.info('timing request', { url })
  https.get(url, function () {
    const duration = Date.now() - start
    log.info('timing request complete', { url, duration })
    res.end(`Duration: ${duration}`)
  })

}).listen(PORT)

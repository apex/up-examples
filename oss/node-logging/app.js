const http = require('http')
const { PORT = 3000 } = process.env

const handlers = {
  text,
  json
}

http.createServer((req, res) => {
  const path = req.url.slice(1)
  const h = handlers[path] || error
  h(req, res)
}).listen(PORT)

/**
 * Text handler.
 */

function text(req, res) {
  console.log('%s %s', req.method, req.url)
  for (let k in req.headers) {
    console.log('  - %s: %s', k, req.headers[k])
  }
  res.end('Hello\n')
}

/**
 * JSON handler.
 */

function json(req, res) {
  log('info', 'performed a request', { method: req.method, path: req.url })
  res.end('Hello\n')
}

/**
 * Default handler.
 */

function error(req, res) {
  res.end('Request /text or /json\n')
}

/**
 * Super simple logger.
 */

function log(level, message, fields = {}) {
  const entry = { level, message, fields }
  console.log(JSON.stringify(entry))
}

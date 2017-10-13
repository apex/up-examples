const http = require('http')
const { PORT = 3000 } = process.env

const handlers = {
  text,
  json,
  stack,
  crash
}

http.createServer((req, res) => {
  const path = req.url.slice(1)
  const h = handlers[path] || error(path)
  h(req, res)
}).listen(PORT)

/**
 * Error handler.
 */

function stack(req, res) {
  try {
    throw new Error('boom!')
  } catch (err) {
    console.error(err.stack)
  }
  res.end(':)\n')
}

/**
 * Crash handler.
 */

function crash(req, res) {
  throw new Error('boom something exploded!')
  res.end(':)\n')
}

/**
 * Text handler.
 */

function text(req, res) {
  console.log('%s %s', req.method, req.url)
  for (let k in req.headers) {
    console.log('  - %s: %s', k, req.headers[k])
  }
  res.end(':)\n')
}

/**
 * JSON handler.
 */

function json(req, res) {
  const { method, url, headers } = req
  log('info', 'performed a request', { method, url, headers })
  res.end(':)\n')
}

/**
 * Default handler.
 */

function error(path) {
  return (req, res) => {
    log('warn', 'invalid path', { path })
    res.write(`Invalid path use the following:\n`)
    res.write('  - /text\n')
    res.write('  - /json\n')
    res.write('  - /stack\n')
    res.write('  - /crash\n')
    res.end()
  }
}

/**
 * Super simple logger.
 */

function log(level, message, fields = {}) {
  const entry = { level, message, fields }
  console.log(JSON.stringify(entry))
}

/**
 * Exception handler.
 */

process.on('uncaughtException', err => {
  log('fatal', 'uncaught exception', {
    error: err.message,
    stack: err.stack
  })

  process.exit(1)
})

const http = require('http')
const { PORT = 3000 } = process.env

// Connect to a database here
// const dbUser = process.env.DB_USER
// const dbPass = process.env.DB_PASS

http.createServer((req, res) => {
  const stage = process.env.UP_STAGE
  const env = process.env.NODE_ENV
  const msg = process.env.MSG
  const name = process.env.NAME
  res.end(`${msg} ${name} from ${stage} (NODE_ENV=${env})\n`)
}).listen(PORT)

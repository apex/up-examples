const http = require('http')
const { PORT = 3000 } = process.env

http.createServer((req, res) => {
  const stage = process.env.UP_STAGE
  const msg = process.env.MSG
  const name = process.env.NAME
  res.end(`${msg} ${name} from ${stage}`)
}).listen(PORT)

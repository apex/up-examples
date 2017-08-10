const http = require('http')

const { PORT = 3000 } = process.env
const env = {}

for (var key in process.env) {
  if (key.indexOf('UP_') == 0) {
    env[key] = process.env[key]
  }
}

http.createServer((req, res) => {
  res.end(JSON.stringify(env, null, 2))
}).listen(PORT)

const express = require('express')
const app = express()

const { PORT = 3000 } = process.env

app.get('/', function(req, res){
  res.send('Hello World from Express!')
})

console.log('listening on %s', PORT)
app.listen(PORT)

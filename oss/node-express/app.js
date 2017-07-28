const express = require('express')
const app = express()

app.get('/', function(req, res){
  res.send('Hello World from Express!')
})

const port = parseInt(process.env.PORT || "3000", 10)
console.log('listening on %s', port)

app.listen(port)

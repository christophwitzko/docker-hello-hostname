var hostname = require('os').hostname()

var express = require('express')
var morgan = require('morgan')

var app = express()

app.use(morgan('combined'))

app.get('/', function (req, res) {
  res.send('Hello World! (' + hostname + ')\n')
})

app.listen(5000, function () {
  console.log('Listening on port 5000 (%s)', hostname)
})

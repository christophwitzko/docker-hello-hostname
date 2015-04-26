import {hostname as osHostname} from 'os'
import express from 'express'
import morgan from 'morgan'

const hostname = osHostname()
const app = express()

app.use(morgan('combined'))

app.get('/', (req, res) => {
  res.send(`Hello World! (${hostname})\n`)
})

app.listen(5000, () => {
  console.log(`Listening on port 5000 (${hostname})`)
})

const Schema = require('graph.ql')
const json = require('koa-json')
const body = require('co-body')
const Koa = require('koa')

const { PORT = 3000 } = process.env

const pets = [
  { id: '0', name: 'Tobi' },
  { id: '1', name: 'Loki' },
  { id: '2', name: 'Jane' }
]

const schema = Schema(`
  type Pet {
    id: Int!
    name: String!
  }

  type Query {
    pet(id: Int!): Pet
  }
`, {
  Query: {
    pet(root, args) {
      return pets[args.id]
    }
  }
})

const app = new Koa

// formatted json output when using ?pretty
app.use(json({ pretty: false, param: 'pretty' }))

// execute the query, expects json input as { query: ... }
// for example try the following request:
// curl -d '{ "query": "{ pet(id: 0) { name }}" }' ':3000/?pretty'
app.use(function *(){
  const { query } = yield body.json(this)
  const { data } = yield schema.query(query)
  this.body = data
})

app.listen(PORT)

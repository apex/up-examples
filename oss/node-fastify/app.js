const fastify = require('fastify')();

const { PORT = 3000 } = process.env;

const schema = {
  schema: {
    response: {
      200: {
        type: 'object',
        properties: {
          hello: {
            type: 'string'
          }
        }
      }
    }
  }
};

fastify
  .get('/', schema, function (req, reply) {
    reply.send({ hello: 'world' });
  });

fastify.listen(PORT, (err) => {
  if (err) {
    throw err;
  }

  console.log(`server listening on ${fastify.server.address().port}`);
});

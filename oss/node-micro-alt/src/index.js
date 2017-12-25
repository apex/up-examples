const micro = require('micro');
const { request } = require('graphql-request');

const { PORT=3000 } = process.env;
const API = 'https://api.graph.cool/simple/v1/movies';

const query = `
  query Movie($title: String!) {
    movie: Movie(title: $title) {
      releaseDate
      actors {
        name
      }
    }
  }
`;

micro(async (req, res) => {
  let title = 'Inception'; // default

  try {
    // Parse `title` from POST requests
    title = (await micro.json(req)).title;
  } catch (err) {
    // micro throws on invalid json
  }

  // Perform & return query
  return await request(API, query, { title });
}).listen(PORT);

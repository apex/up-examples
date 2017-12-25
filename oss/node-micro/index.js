const { json } = require('micro');
const { request } = require('graphql-request');

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

module.exports = async function (req) {
  let title = 'Inception'; // default

  try {
    // Parse `title` from POST requests
    title = (await json(req)).title;
  } catch (err) {
    // micro throws on invalid json
  }

  // Perform query
  return await request(API, query, { title });
}

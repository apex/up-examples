const stage = process.env.UP_STAGE

module.exports = {
  assetPrefix: stage ? `/${stage}` : ''
}

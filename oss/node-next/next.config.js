const upStage = process.env.UP_STAGE;
module.exports = {
  assetPrefix: upStage ? `/${upStage}` : ""
};

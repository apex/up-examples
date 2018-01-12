const polka = require('polka');
const { PORT=3000 } = process.env;

polka()
	.get('/', (req, res) => {
		res.end('Hello World from Polka!');
	})
	.listen(PORT).then(_ => {
		console.log('> listening on %s', PORT);
	});

const express = require('express');
const bodyParser = require('body-parser');
const loginRoute = require('./routes/login');
const PORT = process.env.PORT || 3000;
const rateLimit = require('express-rate-limit');
const app = express();

app.use(bodyParser.json());

app.use('/login', loginRoute);

app.use(
	rateLimit({
		windowMs: 60 * 1000,
		max: 100
	})

);

app.listen(PORT, '0.0.0.0', () => {
	console.log(`Auth service running on port ${PORT}`);
});

app.get('/health', (req, res) => {
	res.status(200).json({ status: 'ok' });
});

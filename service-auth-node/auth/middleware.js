const { verifyToken } = require('./jwt');

module.exports = function authMiddleware(req, res, next) {
	const header = req.headers.authorization;

	if (!header) {
		return res.status(401).json({ error: 'Missing Authorization header' });
	}

	const token = header.split(' ')[1];

	try {
		const payload = verifyToken(token);
		req.user = payload;
		next();
	}	catch (err) {
		return  res.status(401).json({ error: 'Invalid token' });
	}
};

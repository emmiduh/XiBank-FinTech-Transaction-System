const jwt = require('jsonwebtoken');

const SECRET = process.env.JWT_SECRET;

if (!SECRET) {
	throw new Error('JWT_SECRET not set');
}

function generateToken(userId) {
	return jwt.sign({ userId }, SECRET, { expiresIn: '1h' });
}

function verifyToken(token) {
	return jwt.verify(token, SECRET);
}

module.exports = { generateToken, verifyToken };

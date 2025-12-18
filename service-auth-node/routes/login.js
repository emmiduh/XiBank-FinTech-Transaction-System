const express = require('express');
const { generateToken } = require('../auth/jwt');

const router = express.Router();

router.post('/', (req, res) => {
	const { username, password } = req.body;

	if (username && password) {
		const token = generateToken(username);
		return res.json({ token });
	}

	res.status(401).json({ error: 'Invalid credentials' });
});

module.exports = router;

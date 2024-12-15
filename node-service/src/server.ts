// src/index.js
const express = require('express');
const mongoose = require('mongoose');
const cors = require('cors');

const app = express();

app.use(cors());
app.use(express.json());

mongoose.connect(process.env.MONGODB_URI);

const User = mongoose.model('User', {
  id: String,
  name: String,
  email: String
});

app.get('/api/users/:userId', async (req, res) => {
  try {
    const user = await User.findOne({ id: req.params.userId });
    res.json(user);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

const PORT = process.env.PORT || 3001;
app.listen(PORT, () => {
  console.log(`Node.js service running on port ${PORT}`);
});
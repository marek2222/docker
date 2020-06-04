const express = require('express');
const app = express();

app.get('/', (req, res) => res.send({ 'version node': process.version}));

app.listen(8080, () => console.log('Server is up!'));
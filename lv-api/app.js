var express = require('express');
var app = express();
var cors = require('cors');


var swaggerUi = require('swagger-ui-express');
var swaggerDocument = require('./swagger.json');


var LVController = require('./LVController');

app.use(cors());
app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerDocument));
app.use('/', LVController);

module.exports = app;

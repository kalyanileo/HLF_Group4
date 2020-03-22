var app = require('./app');
var port = process.env.PORT || 3000;
var cors = require('cors');


var server = app.listen(port, function() {
  console.log('Express server listening on port ' + port);
});

// var fs = require('fs');
// var http = require('http');
// http.createServer(function(request, response) {
//   var newFile = fs.createWriteStream("readme_copy.md");
//   request.pipe(newFile);
//   request.on('end', function() {
//     response.end('uploaded!');
//   });
// }).listen(8080);
// //curl --upload-file README.md localhost:8080
var request = require('request');
var url = require('url');
app.get('/tweets/:username', function(req, response) {
  var username = req.params.username;
  // route definition
  options = {
    protocol: "http:",
    host: 'api.twitter.com',
    pathname: '/1/statuses/user_timeline.json',
    query: { screen_name: username, count: 10}
    // get the last 10 tweets for screen_name
  }
  var twitterUrl = url.format(options);
  // pipe the request to response
  request(twitterUrl).pipe(response);
});

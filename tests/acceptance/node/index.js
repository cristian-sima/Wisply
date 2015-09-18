var uth = require('./auth.js'),
 webdriverio = require('webdriverio');

var options = {
    desiredCapabilities: {
        browserName: 'firefox'
    }
};

webdriverio
    .remote(options)
    .init()
    .url("http://127.0.0.1:8081")
    .title(function(err, res) {
        console.log('Title was: ' + res.value);
    })
    .end();

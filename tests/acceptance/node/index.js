var uth = require('./auth.js'),
 webdriverio = require('webdriverio');

var options = {
    desiredCapabilities: {
        browserName: 'firefox'
    }
};

/**
 * It shows the introduction to testing
 */
function intro() {
  console.log("Welcome!");
  console.log("Please wait...");
}

/**
 * It shows that the script has been executed
 */
function showHappyEnd() {
  console.log("This is the end of the test. Well done!");
}

function suscces(msg) {
  console.log("[Success] " + msg);
}

function testStaticPages() {
  console.log("[Testing] Static pages");
  webdriverio
      .remote(options)
      .init()
      .url("http://127.0.0.1:8081")
      .title(function(err, res) {
          suscces("The home page is available");
      })
      .url("http://127.0.0.1:8081/contacta")
      .title(function(err, res) {
          suscces("Contact page is available");
      })
      .end();
}

function test() {

  testStaticPages();

}

function start() {
  intro();
  test();
}

start();

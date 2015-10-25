var uth = require('./auth.js'),
 webdriverio = require('webdriverio'),
 client = {};

var options = {
    desiredCapabilities: {
        browserName: 'firefox'
    }
};

client = webdriverio
    .remote(options)
    .init();

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

client.addCommand("checkPageExists", function(url) {
    return this.url(url).then(function(urlResult) {
        return this.isExisting(".wisply-logo").then(function(isExisting) {
            expect(isExisting).toBeFalsy();
        });
    });
});



function testStaticPages() {
  console.log("[Testing] Static pages");
  client
      .url("http://127.0.0.1:8081")
      .checkPageExists("/sfsdfg")
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

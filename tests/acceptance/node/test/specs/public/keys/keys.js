var userModule = require("../../general/user.js"),
  user = userModule.user;

describe('Shortcut Keys', function() {

  it('types ALT + A - Load accessibility bar', function(done) {
    browser
      .url("/")
      .pause(1000)
      .click("#wisply-body")
      .keys(['Alt', 'a'])
      .pause(1500)
      .isExisting('#at-collapse').then(function(isExisting){
          expect(isExisting).toBe(true);
      })
      .call(done);
  });
  it('types ALT + K - Show the key shortcuts', function(done) {
      browser
      .url("/")
      .pause(1000)
      .click("#wisply-body")
      .keys(['Alt', 'k'])
      .pause(1500)
      .isExisting('.modal-title').then(function(isExisting){
          expect(isExisting).toBe(true);
      })
      .getText(".modal-title").then(function(text) {
          expect(text).toContain("Key shortcuts available on this page");
      })
      .call(done);
  });

  it('types ALT + K - Go to home page', function(done) {
      browser
      .url("/contact")
      .pause(1000)
      .click("#wisply-body")
      .keys(['Alt', 'w'])
      .pause(1500)
      .url(function(err, res){
          expect(res.value).toEqual("http://localhost:8081/");
      })
      .call(done);
  });

  it('types ALT + K - Go to contact page', function(done) {
      browser
      .url("/about")
      .pause(1000)
      .click("#wisply-body")
      .keys(['Alt', 'c'])
      .pause(1500)
      .url(function(err, res){
          expect(res.value).toEqual("http://localhost:8081/contact");
      })
      .call(done);
  });


  it('types ALT + K - Go to contact page', function(done) {
      browser
      .url("/about")
      .pause(1000)
      .click("#wisply-body")
      .keys(['Alt', 'c'])
      .pause(1500)
      .url(function(err, res){
          expect(res.value).toEqual("http://localhost:8081/contact");
      })
      .call(done);
  });


  it('connects the account', function(done) {
    browser
    .pause(1000)
    .url("/auth/login")
    .deleteCookie("session")
    .deleteCookie("connection")
    .pause(1000)
    .url("/auth/login")
    .pause(1500)
    .setValue('#login-email', user.email)
    .setValue('#login-password', user.password)
    .submitForm("#login-form")
    .pause(1500)
    .isExisting('#menu-logout-button').then(function(isExisting){
        expect(isExisting).toBe(true);
    })
    .call(done);
  });

  it('types ALT + L - Logs out the current account', function(done) {
      browser
      .url("/about")
      .pause(1000)
      .click("#wisply-body")
      .keys(['Alt', 'l'])
      .pause(1500)
      .isExisting('#menu-logout-button').then(function(isExisting){
          expect(isExisting).toBe(false);
      })
      .call(done);
  });

  it('types CTRL + Space - Focuses the search field', function(done) {
      browser
      .url("/about")
      .pause(1000)
      .click("#wisply-body")
      .pause(1000)
      .keys('Ctrl').keys('U+E00D').keys('a').keys('NULL')
      .pause(1500)
      .elementActive().then(function(element){
          element = true;
          expect(element).toBe(true);
      })
      .call(done);
  });
});

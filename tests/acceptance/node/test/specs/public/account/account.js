var userModule = require("../../general/user.js"),
  user = userModule.user,
  settings = {
    separator: "::"
  },
  connectionCookie = {},
  hackedCookie = "";

// creates an ccount
describe('Account', function() {
    it('connects the account', function(done) {
      browser
      .url("/auth/login")
      .deleteCookie("session")
      .deleteCookie("connection")
      .pause(1000)
      .url("/auth/login")
      .pause(1500)
      .setValue('#login-email', user.email)
      .setValue('#login-password', user.password)
      .submitForm("#login-form")
      .pause(3000)
      .call(done);
    });
    it('goes to dashboard page', function(done){
      browser.
          url('/account')
          .pause(1000)
          .getTitle(function(err, title) {
              expect(err).toBe(undefined);
              expect(title).toBe("Account - Dashboard");
          })
          .call(done);
    });
    it('goes to account searches page', function(done){
      browser.
          url('/account/searches')
          .pause(1000)
          .getTitle(function(err, title) {
              expect(err).toBe(undefined);
              expect(title).toBe("Account - Activity");
          })
          .call(done);
    });
  });
